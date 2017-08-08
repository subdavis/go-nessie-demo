package nessie

import (
	"encoding/json"
	"io/ioutil"
	_ "bytes"
	"net/http"
)

type (

  Merchant struct {
  	Id          string    `json:"_id"`
  	// Please vadidate data types server-side...  There are "string" geocoordinates....
  	Coordinates map[string]interface{} `json:"geocode"` 
  	Name        string    `json:"name"`
  	Address     Address   `json:"address"`
  }
  MerchantResponse struct {
  	Results []Merchant    `json:"results"`
  }
)

func (c Client) GetMerchants () []Merchant {
	var mR MerchantResponse
	url := BaseURL + "/enterprise/merchants" + c.EncodeParams(nil)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &mR)
	if err != nil {
		panic(err)
	} 
	return mR.Results
}