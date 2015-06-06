package coinbase

import (
	"fmt"
	"log"
	"os"
	"testing"

	"com.teddywing/new-house-on-the-block/vendor/_nuts/github.com/stretchr/testify/assert"
)

// Initialize the client without mock mode enabled on rpc.
// All calls hit the coinbase API and tests focus on checking
// format of the response and validity of sent requests
func initClient() Client {
	return ApiKeyClient(os.Getenv("COINBASE_KEY"), os.Getenv("COINBASE_SECRET"))
}

// About Endpoint Tests:
//All Endpoint Tests actually call the Coinbase API and check the return values
// with type assertions. This was done because of the varying specific values
// returned depending on the API Key and Secret pair used when running the tests.
// Endpoint tests do not include tests that could be run an arbitrary amount of times
// i.e buy, sell, etc...

func TestEndpointGetBalance(t *testing.T) {
	c := initClient()
	amount, err := c.GetBalance()
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, 1.1, amount)
}

func TestEndpointGetReceiveAddress(t *testing.T) {
	t.Skip("Skipping GetReceiveAddressEndpoint in order not to create excessive amounts of receive addresses during testing.")
	c := initClient()
	params := &AddressParams{
		CallbackUrl: "http://www.wealthlift.com",
		Label:       "My Test Address",
	}
	address, err := c.GenerateReceiveAddress(params)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", address)
}

func TestEndpointGetAllAddresses(t *testing.T) {
	c := initClient()
	params := &AddressesParams{
		Page:  1,
		Limit: 5,
	}
	addresses, err := c.GetAllAddresses(params)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", addresses.Addresses[0].CreatedAt)
	assert.IsType(t, "string", addresses.Addresses[0].Address)
}

func TestEndpointCreateButton(t *testing.T) {
	c := initClient()
	params := &Button{
		Name:             "test",
		Type:             "buy_now",
		Subscription:     false,
		PriceString:      "1.23",
		PriceCurrencyIso: "USD",
		Custom:           "Order123",
		CallbackUrl:      "http://www.example.com/my_custom_button_callback",
		Description:      "Sample Description",
		Style:            "custom_large",
		IncludeEmail:     true,
	}
	data, err := c.CreateButton(params)
	if err != nil {
		if fmt.Sprint(err) != "You have not filled out your merchant profile. Please enter your information in the Profile section. in CreateButton()" {
			log.Fatal(err)
		}
		t.Skip("Skip this test since user hasn't filled out their merchant profile yet.")
	}
	assert.IsType(t, "string", data.EmbedHtml)
	assert.IsType(t, "string", data.Type)
}

func TestEndpointGetCurrencies(t *testing.T) {
	c := initClient()
	data, err := c.GetCurrencies()
	if err != nil {
		log.Fatal()
	}
	assert.IsType(t, "string", data[0].Name)
}

func TestEndpointGetExchangeRates(t *testing.T) {
	c := initClient()
	data, err := c.GetExchangeRates()
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data["btc_to_usd"])
}

func TestEndpointGetExchangeRate(t *testing.T) {
	c := initClient()
	data, err := c.GetExchangeRate("btc", "usd")
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, 0.0, data)
}

func TestEndpointGetTransactions(t *testing.T) {
	c := initClient()
	data, err := c.GetTransactions(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, int64(1), data.TotalCount)
	assert.IsType(t, "string", data.Transactions[0].Hsh)
}

func TestEndpointGetBuyPrice(t *testing.T) {
	c := initClient()
	data, err := c.GetBuyPrice(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data.Subtotal.Currency)
	assert.IsType(t, "string", data.Total.Amount)
}

func TestEndpointGetSellPrice(t *testing.T) {
	c := initClient()
	data, err := c.GetSellPrice(1)
	if err != nil {
		log.Fatal(err)
	}
	assert.IsType(t, "string", data.Subtotal.Currency)
	assert.IsType(t, "string", data.Total.Amount)
}

func TestEndpointGetTransaction(t *testing.T) {
	c := initClient()
	_, err := c.GetTransaction("5446968682a19ab940000004")
	if err != nil {
		assert.IsType(t, "string", err.Error())
	}
}
