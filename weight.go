package furrytail

import (
	"errors"
	"strconv"
)

//GetCurrentWeight returns weight of food in the bowl
func (a *Account) GetCurrentWeight(deviceID int) (grams int, err error) {
	res, err := a.client.R().SetPathParams(map[string]string{
		"deviceId": strconv.Itoa(deviceID),
	}).SetResult(&weightResponse{}).Put("/provider-feeder/feeder/weigh/{deviceId}")
	if err != nil {
		return -1, err
	}
	if res.IsError() {
		return -1, errors.New(res.Error().(*apiError).Message)
	}
	return res.Result().(*weightResponse).Data, nil
}

type weightResponse struct {
	apiResponse
	Data int `json:"data"`
}
