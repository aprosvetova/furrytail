package furrytail

import "errors"

//SetManualFeeding enables or disables manual feeding via physical button
func (a *Account) SetManualFeeding(deviceID int, allowed bool) error {
	value := 0
	if allowed {
		value = 1
	}
	res, err := a.client.R().SetResult(&boolResponse{}).SetBody(optionRequest{
		DeviceID: deviceID,
		Value:    value,
	}).Put("/provider-feeder/feeder/feed/manual/status")
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

//SetLight enables or disables circle LED indication
func (a *Account) SetLight(deviceID int, enabled bool) error {
	value := 0
	if enabled {
		value = 1
	}
	res, err := a.client.R().SetResult(&boolResponse{}).SetBody(optionRequest{
		DeviceID: deviceID,
		Value:    value,
	}).Put("/provider-feeder/feeder/lighting/status")
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

type optionRequest struct {
	DeviceID int `json:"deviceId"`
	Value    int `json:"value"`
}
