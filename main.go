package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/teddywing/new-house-on-the-block/purchase"
)

func sendMoneyToSeller() error {
	transaction_id, err := purchase.SendMoney(os.Getenv("COINBASE_KEY"),
		os.Getenv("COINBASE_SECRET"),
		"n2Qd6da1jiFgij5SSncFKh7MoFN74GdUxv",
		"0.0001")
	if err != nil {
		return err
	} else {
		fmt.Println(transaction_id)
		return nil
	}
}

func sendTokenToBuyer() error {
	transaction_id, err := purchase.SendMoney(os.Getenv("SELLER_COINBASE_KEY"),
		os.Getenv("SELLER_COINBASE_SECRET"),
		"mqy3kT6aFHymTcvmdwZLKq1Svo2m6sUtzH",
		"0.0001")
	if err != nil {
		return err
	} else {
		fmt.Println(transaction_id)
		return nil
	}
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/buy/", func(w http.ResponseWriter, r *http.Request) {
		var err error

		err = sendMoneyToSeller()
		if err != nil {
			log.Println(err)
		}

		err = sendTokenToBuyer()
		if err != nil {
			log.Println(err)
		}
		
		http.Redirect(w, r, "http://testsite.perfectburpee.com/new-house-on-the-block-page6/", 302)
	})

	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
