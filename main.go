package main

import (
	"com.teddywing/new-house-on-the-block/vendor/_nuts/github.com/fabioberger/coinbase-go"
	"fmt"
	"log"
	"net/http"
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
	// transaction_id, err := sendMoney("TODO", "n2Qd6da1jiFgij5SSncFKh7MoFN74GdUxv", "0.0001")
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println(transaction_id)
	// }

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
