package furrytail

import (
	"github.com/go-resty/resty"
	"time"
)

//Account is a main object containing all methods for working with feeder devices
type Account struct {
	client *resty.Client
	loc time.Location
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
