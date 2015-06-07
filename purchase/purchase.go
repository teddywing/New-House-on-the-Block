package purchase

import (
	"github.com/teddywing/new-house-on-the-block/vendor/_nuts/github.com/fabioberger/coinbase-go"
)

func SendMoney(from_key string, from_secret string, to string, amount string) (transaction_id string, err error) {
	c := coinbase.ApiKeyClientSandbox(from_key, from_secret)

	params := &coinbase.TransactionParams{
		To:     to,
		Amount: amount,
		Notes:  "You just bought a house",
	}

	confirmation, err := c.SendMoney(params)
	if err != nil {
		return "", err
	} else {
		return confirmation.Transaction.Id, nil
	}
}
