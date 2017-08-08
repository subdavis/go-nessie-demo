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

	customer := CreateCustomer(client)
	CreateAccount(client, customer)
	accounts := GetAccounts(client)
	ShowNearbyATMs(client, 3)
	CreateBill(client, accounts[0])
	merchants := GetMerchants(client)
	DoPurchase(client, accounts[0], merchants[0])
	CreateDeposit(client, accounts[0])
	CleanUp(client)
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func GetAccounts(c nessie.Client) []nessie.Account {
	var accounts []nessie.Account = c.GetAccounts()

	fmt.Println("Your accounts are....")
	for _, a := range accounts {
		fmt.Println("  ", a.Nickname, " ", a.Id)
	}
	fmt.Println()

	if len(accounts) == 0 {
		panic("No accounts ID to use for bill creation.  Please create an account related to your key or choose another key.")
	}

	return accounts
}

func CreateAccount(c nessie.Client, customer nessie.NessieObject) {
	fmt.Println("Creating Account for Customer with ID ", customer.Id)
	accountData := []byte(`{
	  "type": "Checking",
	  "nickname": "Brandon's 360 Checking",
	  "rewards": 0,
	  "balance": 0,
	  "account_number": "1111222233334444"
	}`)
	fmt.Println(string(accountData))
	acct := c.CreateAccount(customer.Id, accountData)
	fmt.Println("  Created Account? ", acct)
	fmt.Println()
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
func CreateCustomer(c nessie.Client) nessie.NessieObject{
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
	cust := c.CreateCustomer(customerData)
	fmt.Println("  Created Customer? ", cust.Id)
	fmt.Println()
	return cust
}

//Uses at least one endpoint for customers, accounts and bills, two of which must be a POST request.
func CreateBill(c nessie.Client, account nessie.Account){
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
	status := c.CreateBill(account.Id, billData)
	fmt.Println("  Created Bill? ", status)
	fmt.Println()
}

//Uses one purchase endpoint
func DoPurchase(c nessie.Client, account nessie.Account, merchant nessie.Merchant){
	purchaseData := []byte(`  {
	  "merchant_id": "`+merchant.Id+`",
	  "medium": "balance",
	  "purchase_date": "2017-07-31",
	  "amount": 4.55,
	  "description": "`+merchant.Name+`"
	}`)
	fmt.Println("Creating Purchase for " + account.Nickname + " with details...")
	fmt.Println(string(purchaseData))
	status := c.CreatePurchase(account.Id, purchaseData)
	fmt.Println("  Created Purchase? ", status)
	fmt.Println()
}

//Uses one money movement endpoint (deposit, withdrawal, transfer) that is NOT a GET request
func CreateDeposit(c nessie.Client, account nessie.Account) {
	depositData := []byte(`{
	  "medium": "balance",
	  "transaction_date": "2017-08-31",
	  "amount": 100.0,
	  "description": "Test Deposit on 8.31.17"
	}`)
	fmt.Println("Creating deposit for " + account.Nickname + " with details...")
	fmt.Println(string(depositData))
	status := c.CreateDeposit(account.Id, depositData)
	fmt.Println("  Created Deposit? ", status)
	fmt.Println()
}

//Uses one enterprise endpoint
func GetMerchants(c nessie.Client) []nessie.Merchant {
	fmt.Println("Getting list of merchants...")
	merchants := c.GetMerchants()
	fmt.Println("  There are " , len(merchants), " merchants.")
	if len(merchants) == 0 {
		panic("There are no merchants.  Something must be wrong...")
	}
	fmt.Println()
	return merchants
}

//Use the DELETE /data endpoint to delete a data entity(Accounts, Customers, etc) of your choice
func CleanUp(c nessie.Client) {
	fmt.Println("Cleaning up a few things.")
	status := c.DeleteData("Deposits")
	fmt.Println("  Removing Deposits? ", status)
	status = c.DeleteData("Customers")
	fmt.Println("  Remove customers? ", status)
	status =c.DeleteData("Accounts")
	fmt.Println("  Removing Accounts? ", status)
}