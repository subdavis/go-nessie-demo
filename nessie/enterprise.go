package nessie

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	_ "bytes"
	"strings"
	"net/http"
)

type (

  Merchant struct {
  	Id          string    `json:"_id"`
  	Coordinates Location  `json:"geocode"`
  	Name        string    `json:"name"`
  	Address     Address   `json:"address"`
  }
  MerchantResponse struct {
  	Results []Merchant    `json:"results"`
  }
)

func (c Client) CountMerchants () int {
	url := BaseURL + "/enterprise/merchants" + c.EncodeParams(nil)
	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	dec := json.NewDecoder(strings.NewReader(string(body)))
	var m map[string]interface{}
	err = dec.Decode(&m)
	_ = err
	return resp.StatusCode 
}