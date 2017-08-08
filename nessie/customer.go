package nessie

import (
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func (c Client) CreateCustomer (CustomerData []byte) NessieObject {
	url := BaseURL + "/customers" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(CustomerData))
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var customer ResponseClass
	err = json.Unmarshal(body, &customer)
	if err != nil {
		panic(err)
	}
	return customer.ObjectCreated
}

func (c Client) CreateAccount (CustomerId string, AccountData []byte) int {
	url := BaseURL + "/customers/" + CustomerId + "/accounts" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(AccountData))
	if err != nil {
		panic(err)
	}
	return resp.StatusCode
}

func (c Client) DeleteData (dataType string) int {
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
	return resp.StatusCode
}