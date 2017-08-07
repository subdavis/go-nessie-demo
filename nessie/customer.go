package nessie

import (
	_ "fmt"
	"bytes"
	"net/http"
)

func (c Client) CreateCustomer (CustomerData []byte) bool {
	url := BaseURL + "/customers" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(CustomerData))
	if err != nil {
		panic(err)
	}
	return resp.StatusCode == 201
}

func (c Client) DeleteData (dataType string) bool {
	url := BaseURL + "/data" + c.EncodeParams(map[string]string{"type":dataType})
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode == 204
}