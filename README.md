New House on the Block
======================

A project created at AngelHack Boston 2015.06. The idea is to use the Bitcoin
block chain to facilitate real estate transactions.

It works by assuming that we can associate a token amount of Bitcoin with a
property and call that token the title. Thus, the owner of the token becomes
the owner of the property. The legal implications of this assumption are not
addressed. This is more of a thought experiment.

Continuing under this assumption, it's possible to streamline the transaction
of the title and payment using Bitcoin.

The buyer initiates a multi-signature transaction with the seller, in which the
payment `P` will be sent to the seller and the title/token `T` will be sent to
the buyer. After the buyer signs this transaction, it is sent to the seller. At
this point, the transaction is not complete, and no money has changed hands.
Once the seller signs and the transaction is validated, the payment and token
change hands.

	Buyer creates:                                        
														  
	Transaction                                           
	+-------------------+------------------+              
	|                   |                  |              
	| P BTC  ->  Seller | Buyer signature  |              
	|                   |                  |              
	+-------------------+------------------+              
	|                   |                  |              
	| Buyer  <-  T BTC  |                  |              
	|                   |                  |              
	+-------------------+------------------+              
														  
						+                                 
						|                                 
						v                                 
														  
	Transaction sent to Seller                            
														  
	Seller adds signature                                 
	+-------------------+------------------+              
	|                   |                  |              
	| P BTC  ->  Seller | Buyer signature  |              
	|                   |                  |              
	+-------------------+------------------+              
	|                   |                  |              
	| Buyer  <-  T BTC  | Seller signature |              
	|                   |                  |              
	+-------------------+------------------+              
														  
	Transaction is then validated                         
	and completes                                         

Many portions of the code are hard-coded to facilitate easy demoing at the
hackathon. It requires 2 [Coinbase](https://www.coinbase.com/) accounts. I took
advantage of the [Coinbase Sandbox](https://sandbox.coinbase.com/) in order to
save myself the trouble of dealing with real Bitcoins.

As of this writing, Coinbase doesn't expose a way to perform “multisig”
transactions in their sandbox environment, so this demo emulates the process by
initiating 2 different transactions instead of 1. This doesn't provide any of
the trust from the single-transaction model described above, but it works well
enough for the purpose of a demo.

The code creates an HTTP server and provides an endpoint (`/buy/`) that will
perform 2 transactions, sending Bitcoin from buyer to seller and “token” (in
this case more Bitcoin) from seller to buyer.


## Running
1. `$ cp .env.sample .env`

	Add to `.env` the API keys and secrets from 2 Coinbase accounts.

2. `$ source .env`

3. Replace the Bitcoin addresses in [main.go](main.go) with the addresses of your
2 accounts.

4. `$ go build`

5. `$ ./new-house-on-the-block`

6. Visit `http://localhost:3000/buy/` to initiate the transaction.


## License
Licensed under the MIT License. See the included LICENSE file.
