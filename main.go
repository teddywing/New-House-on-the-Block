package main

import (
	"com.teddywing/new-house-on-the-block/vendor/_nuts/github.com/fabioberger/coinbase-go"
	"fmt"
	"log"
	"os"
)

func sendMoney(from string, to string, amount string) (transaction_id string, err error) {
	c := coinbase.ApiKeyClientSandbox(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))

	params := &coinbase.TransactionParams{
		To:     to,
		Amount: amount,
		Notes:  "You just bought a house",
	}

	confirmation, err := c.SendMoney(params)
	if err != nil {
		return "", err
	} else {
		fmt.Println(confirmation.Transaction.Status)
		fmt.Println(confirmation.Transaction.Id)

		return confirmation.Transaction.Id, nil
	}
}

func main() {
	transaction_id, err := sendMoney("TODO", "n2Qd6da1jiFgij5SSncFKh7MoFN74GdUxv", "0.0001")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(transaction_id)
	}
}
