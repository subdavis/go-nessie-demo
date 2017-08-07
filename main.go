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

	account_ids := ShowAccounts(client)
	ShowNearbyATMs(client, 3)
	CreateCustomer(client)
	CreateBill(client, account_ids[0])
	DoPurchase(client, account_ids[0])
	CleanUp(client)
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func ShowAccounts(c nessie.Client) []string {
	var accountIDs []string = make([]string, 0)
	var accounts []nessie.Account = c.GetAccounts()

	fmt.Println("Your accounts are....")
	for _, a := range accounts {
		accountIDs = append(accountIDs, a.Id)
		fmt.Println("  ", a.Nickname)
	}
	fmt.Println()

	if len(accountIDs) == 0 {
		panic("No account ID to use for bill creation.")
	}

	return accountIDs
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
  customerData := []byte(`  {
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
	fmt.Println("  Created Customer? ", status)
	fmt.Println()
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func CreateBill(c nessie.Client, accountId string){
	fmt.Println("Creating Spotify Monthly Bill with details...")
	billData := []byte(`  {
	  "status": "pending",
	  "payee": "Spotify",
	  "nickname": "Spotify Monthly Bill",
	  "payment_date": "2017-08-31",
	  "recurring_date": 31,
	  "payment_amount": 9.99
	}`)
	fmt.Println(string(billData))
	status := c.CreateBill(accountId, billData)
	fmt.Println("  Created Bill? ", status)
	fmt.Println()
}

//Uses one purchase endpoint
func DoPurchase(c nessie.Client, accountId string){
	purchaseData := []byte(`  {
	  "merchant_id": "57cf75cea73e494d8675ec49",
	  "medium": "balance",
	  "purchase_date": "2017-07-31",
	  "amount": 4.55,
	  "description": "Dunkin Donuts Chapel Hill, NC"
	}`)
	fmt.Println("Creating Purchase for " + accountId + " with details...")
	fmt.Println(string(purchaseData))
	status := c.CreatePurchase(accountId, purchaseData)
	fmt.Println("  Created Purchase? ", status)
	fmt.Println()
}

//Uses one money movement endpoint (deposit, withdrawal, transfer) that is NOT a GET request
func CreateTransfer(c nessie.Client) {
	fmt.Println()
}

//Uses one enterprise endpoint
func CountAllAccounts(c nessie.Client) {
	fmt.Println()
}

//Use the DELETE /data endpoint to delete a data entity(Accounts, Customers, etc) of your choice
func CleanUp(c nessie.Client) {
	fmt.Println("Cleaning up a few things.")
	status := c.DeleteData("Customers")
	fmt.Println("  Remove customers? ", status)
	status =c.DeleteData("Purchases")
	fmt.Println("  Removing Purchases? ", status)
}