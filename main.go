package main

import (
	"os"
	"fmt"
	"github.com/subdavis/go-nessie-demo/nessie"
)

func main(){

	client, err := nessie.NewClient(os.Getenv("NESSIE_KEY"))
	if err != nil {
		panic(err)
	}
	fmt.Println("You're using key:",client.Key)

	ShowAccountsCreateBills(client)
	//ShowNearbyATMs(client, 3)
	//CreateCustomer(client)
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func ShowAccountsCreateBills(c nessie.Client) {
	var accountID string
	fmt.Println("Your accounts are....")
	var accounts []nessie.Account = c.GetAccounts()
	for _, a := range accounts {
		accountID = a.Id
		fmt.Println("  ", a.Nickname)
	}
	fmt.Println()

	if accountID == "" {
		panic("No account ID to use for bill creation.")
	}
	fmt.Println("Creating Spotify Monthly Bill with details...")
	billData := []byte(`{
	  "status": "pending",
	  "payee": "Spotify",
	  "nickname": "Spotify Monthly Bill",
	  "payment_date": "2017-08-31",
	  "recurring_date": 31,
	  "payment_amount": 9.99
	}`)
	fmt.Println(string(billData))
	status := c.CreateBill(accountID, billData)
	fmt.Println("Created Bill? ", status)
}

// Uses the GET /atms endpoint. Please note it is paginated, and your submission must query ATMs multiple times using the paging object.
func ShowNearbyATMs(c nessie.Client, radius int64) {
  fmt.Println("ATMS within a ",radius, " mile radius to your location...")
	var atms []nessie.ATM = c.GetNearbyATMs(radius)
	for _, a := range atms {
		fmt.Println("  ", a.Name)
	}
	fmt.Println()
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func CreateCustomer(c nessie.Client) {
  customerData := []byte(`{
	  "first_name": "Brandon",
	  "last_name": "Davis",
	  "address": {
	    "street_number": "1680",
	    "street_name": "Capital One Dr",
	    "city": "McLean",
	    "state": "VA",
	    "zip": "22102"
	  }
	}`)
	fmt.Println("Creating Customer with details...")
	fmt.Println(string(customerData))
	var status bool = c.CreateCustomer(customerData)
	fmt.Println("Created Customer? ", status)
}

//Uses one purchase endpoint
func DoPurchase(c nessie.Client){

}

//Uses one money movement endpoint (deposit, withdrawal, transfer) that is NOT a GET request
func CreateTransfer(c nessie.Client) {

}

//Uses one enterprise endpoint
func CountAllAccounts(c nessie.Client) {

}

//Use the DELETE /data endpoint to delete a data entity(Accounts, Customers, etc) of your choice
func CleanUp(c nessie.Client) {

}