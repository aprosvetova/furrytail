package furrytail

import (
	"errors"
	"strconv"
)

//ListDevices returns all available feeder devices on the account
func (a *Account) ListDevices() (devices []Device, err error) {
	res, err := a.client.R().SetResult(&deviceListResponse{}).Get("/provider-device/user/device/list")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New("api error")
	}
	for _, d := range res.Result().(*deviceListResponse).Data {
		devices = append(devices, convertDevice(d))
	}
	return
}

//GetDevice returns latest information about specific feeder device
func (a *Account) GetDevice(deviceID int) (device *Device, err error) {
	res, err := a.client.R().SetPathParams(map[string]string{
		"deviceId": strconv.Itoa(deviceID),
	}).SetResult(&productResponse{}).Get("/provider-feeder/feeder/info/{deviceId}")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, errors.New(res.Error().(*apiError).Message)
	}
	return convertProduct(res.Result().(*productResponse).Data), nil
}

type deviceListResponse struct {
	apiResponse
	Data []apiDevice
}

type productResponse struct {
	apiResponse
	Data apiProduct
}

type apiDevice struct {
	ID      int        `json:"id"`
	Name    string     `json:"name"`
	Online  int        `json:"onlineStatus"`
	Product apiProduct `json:"product"`
}

type apiProduct struct {
	BucketStatus     int    `json:"bucketFoodStatus"`
	CalibrationTimes int    `json:"calibrationTimes"`
	DcMotorStatus    int    `json:"dcMotorStatus"`
	Electricity      int    `json:"electricity"`
	LightStatus      int    `json:"lightStatus"`
	ManualFeed       int    `json:"manualFeed"`
	Ssid             string `json:"ssid"`
	Status           int    `json:"status"`
	StepMotorStatus  int    `json:"stepMotorStatus"`
	Version          string `json:"version"`
	Device           struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Online int    `json:"onlineStatus"`
	} `json:"device"`
}

//Device is an object representing feeder with it's status
type Device struct {
	ID                   int
	Name                 string
	Online               bool
	BucketFull           bool
	DcMotorProblem       bool
	StepMotorProblem     bool
	BatteryLevelHours    int
	Light                bool
	ManualFeedingAllowed bool
	Ssid                 string
	Version              string
}

func convertDevice(input apiDevice) Device {
	return Device{
		ID:                   input.ID,
		Name:                 input.Name,
		Online:               input.Online == 1,
		BucketFull:           input.Product.BucketStatus == 1,
		DcMotorProblem:       input.Product.DcMotorStatus == 0,
		StepMotorProblem:     input.Product.StepMotorStatus == 0,
		BatteryLevelHours:    input.Product.Electricity,
		Light:                input.Product.LightStatus == 1,
		ManualFeedingAllowed: input.Product.ManualFeed == 1,
		Ssid:                 input.Product.Ssid,
		Version:              input.Product.Version,
	}
}

func convertProduct(input apiProduct) *Device {
	return &Device{
		ID:                   input.Device.ID,
		Name:                 input.Device.Name,
		Online:               input.Device.Online == 1,
		BucketFull:           input.BucketStatus == 1,
		DcMotorProblem:       input.DcMotorStatus == 0,
		StepMotorProblem:     input.StepMotorStatus == 0,
		BatteryLevelHours:    input.Electricity,
		Light:                input.LightStatus == 1,
		ManualFeedingAllowed: input.ManualFeed == 1,
		Ssid:                 input.Ssid,
		Version:              input.Version,
	}
}
