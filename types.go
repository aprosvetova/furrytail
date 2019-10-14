package furrytail

import "github.com/go-resty/resty"

//Account is a main object containing all methods for working with feeder devices
type Account struct {
	client *resty.Client
}

type apiResponse struct {
	Code   string `json:"code"`
	Status bool   `json:"status"`
}

type apiError struct {
	apiResponse
	Message string `json:"message"`
}

type boolResponse struct {
	apiResponse
	Data bool `json:"data"`
}
