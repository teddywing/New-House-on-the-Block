package main

import (
	"com.teddywing/new-house-on-the-block/vendor/_nuts/github.com/fabioberger/coinbase-go"
	"fmt"
	"log"
	"os"
)

func main() {
	c := coinbase.ApiKeyClientSandbox(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))

	balance, err := c.GetBalance()
	if err != nil {
		log.Print(err)
	}
	fmt.Printf("Balance is %f BTC\n", balance)
}
