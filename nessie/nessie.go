package nessie

import (
	"errors"
	_ "fmt"
)

type (

	Location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}

	Address struct {
		City       string `json:"city"`
		StreetNum  string `json:"street_number"`
		Zip        string `json:"zip"`
		State      string `json:"state"`
		StreetName string `json:"street_name"`
	}
)

var (

	BaseURL = "http://api.reimaginebanking.com"
)

type (
	Client struct {
		Key string
	}
)

func NewClient (api_key string) (Client, error) {
	var c Client
	if api_key == "" {
		return c, errors.New("Cannot have empty api_key")
	}
	c = Client{
		Key: api_key,
	}
	return c, nil
}

func (c Client) EncodeParams (params map[string]string) string {
	var paramString string = "?"
	//TODO: urlencode
	for k, v := range params {
		paramString += k + "=" + v + "&"
	}
	paramString += "key=" + c.Key
	return paramString
}