package nessie

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"bytes"
	"strings"
	"net/http"
)

func (c Client) CreateCustomer (CustomerData []byte) bool {
	url := BaseURL + "/customers" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(CustomerData))
	return resp.StatusCode == 201
}