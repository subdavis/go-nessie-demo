package nessie

import (
	"encoding/json"
	_ "fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

type (

	Account struct {
		Id            string `json:"_id"`
		AccountType   string `json:"Type"`
		Nickname      string `json:"nickname"`
		Rewards       int    `json:"rewards"`
		Balance       int    `json:"blanace"`
		AccountNumber string `json:"account_number"`
		CustomerId    string `json:"customer_id"`
	}
)

func (c *Client) GetAccounts() []Account {
	url := BaseURL + "/accounts" + c.EncodeParams(nil)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var accounts []Account
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		panic(err)
	}
	return accounts
}

func (c Client) CreateBill (AccountId string, BillData []byte) int {
	url := BaseURL + "/accounts/" + AccountId + "/bills" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(BillData))
	if err != nil {
		panic(err)
	}
	return resp.StatusCode
}

func (c Client) CreatePurchase (AccountId string, PurchaseData []byte) int {
	url := BaseURL + "/accounts/" + AccountId + "/purchases" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(PurchaseData))
	if err != nil {
		panic(err)
	}
	return resp.StatusCode
}

func (c Client) CreateDeposit (AccountId string, DepositData []byte) int {
	url := BaseURL + "/accounts/" + AccountId + "/deposits" + c.EncodeParams(nil)
	resp, err := http.Post(url, "application/json", bytes.NewReader(DepositData))
	if err != nil {
		panic(err)
	}
	return resp.StatusCode
}