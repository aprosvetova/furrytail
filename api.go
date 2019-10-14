package furrytail

import (
	"errors"
	"github.com/go-resty/resty"
)

//New creates new Furrytail Account instance and checks if the token is correct
func New(token string) (*Account, error) {
	client := resty.New()
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		req.SetHeader("Authorization", token)
		return nil
	})
	client.OnAfterResponse(func(c *resty.Client, res *resty.Response) error {
		if !res.IsError() {
			return nil
		}
		err, ok := res.Error().(*apiError)
		if !ok {
			return nil
		}
		err.Message = translateError(err.Message)
		return nil
	})
	client.SetHostURL("https://api.furrytail-pet.com")
	client.SetHeader("platformtype", "3")
	client.SetHeader("appversion", "1.0.8")
	client.SetError(&apiError{})

	acc := Account{
		client: client,
	}

	err := acc.refreshToken()
	if err != nil {
		return nil, err
	}

	return &acc, nil
}

func (a *Account) refreshToken() error {
	res, err := a.client.R().Put("/provider-user/user/token/refresh")
	if res.IsError() {
		return errors.New("wrong token")
	}
	return err
}
