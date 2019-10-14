package furrytail

import (
	"errors"
)

//Feed dispenses food immediately. Copies parameter is portions count (1 portion = ~6 grams)
func (a *Account) Feed(deviceID int, copies int) error {
	res, err := a.client.R().SetResult(&boolResponse{}).SetBody(feedRequest{
		DeviceID: deviceID,
		Copies:   copies,
	}).Post("/provider-feeder/feed/log/item/manual")
	if err != nil {
		return err
	}
	if res.IsError() {
		return errors.New(res.Error().(*apiError).Message)
	}
	if !res.Result().(*boolResponse).Data {
		return errors.New("unknown error")
	}
	return nil
}

type feedRequest struct {
	DeviceID int `json:"deviceId"`
	Copies   int `json:"copies"`
}
