package nessie

import (
	"encoding/json"
	"io"
	"log"
	"strings"
	"io/ioutil"
	"strconv"
	"net/http"
)

type (
	
	ATM struct {
		Id            string   `json:"_id"`
		Name          string   `json:"name"`
		Languages     []string `json:"language_list"`
		Coordinates   Location `json:"geocode"`
		Hours         []string `json:"hours"`
		Accessibility bool     `json:"accessibility"`
		AmountLeft    int      `json:"amount_left"`
	}

	ATMPage struct {
		Paging PageMeta `json:"paging"`
		Data   []ATM    `json:"data"`
	}
)

func GetLatLng () (float64, float64) {
	currentGPSurl := "https://freegeoip.net/json/"
	resp, err := http.Get(currentGPSurl)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	var m map[string]interface{}
	err = dec.Decode(&m)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return m["latitude"].(float64), m["longitude"].(float64)
}

func (c *Client) GetNearbyATMs(radius int64) []ATM {

	var atms []ATM = make([]ATM, 0)
	var page ATMPage

	params := make(map[string]string)
	lat, lng := GetLatLng()
	params["lat"] = strconv.FormatFloat(lat, 'f', 4, 64)
	params["lng"] = strconv.FormatFloat(lng, 'f', 4, 64)
	params["rad"] = strconv.FormatInt(radius, 10) // 10 mile radius

	url := BaseURL + "/atms" + c.EncodeParams(params)

	for {

	  page = *new(ATMPage) // zero the memory

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(body, &page)
		if err != nil {
			panic(err)
		}

		atms = append(atms, page.Data...)
		next := page.Paging.Next

		if next != "" {
			url = BaseURL + next
		} else {
			break
		}
	}
	return atms
}
