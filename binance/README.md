## Import Package

```go
import "github.com/aaabigfish/goex/binance"
```

## Public API Ticker Example

```go
    ticker, err := binance.NewBinance().FetchTicker(context.Background(), currency.NewPair(currency.BTC, currency.USDT), asset.Spot)
    if err != nil {
        // Handle error
    }
    fmt.Println(ticker.Last)
```

## Private API Submit Order Example

```go
    b := binance.NewBinance()

	// Set default keys
	b.API.SetKey("your_key")
	b.API.SetSecret("your_secret")
	b.API.SetClientID("your_clientid")
	b.API.SetPEMKey("your_PEM_key")
	b.API.SetSubAccount("your_specific_subaccount")

	// Set client/strategy/subsystem specific credentials that will override
	// default credentials.
	// Make a standard context and add credentials to it by using exchange
	// package helper function DeployCredentialsToContext
	ctx := context.Background()
	ctx = account.DeployCredentialsToContext(ctx, &account.Credentials{
		Key:        "your_key",
		Secret:     "your_secret",
		ClientID:   "your_clientid",
		PEMKey:     "your_PEM_key",
		SubAccount: "your_specific_subaccount",
	})

	o := &order.Submit{
		Exchange:  b.Name, // or method GetName() if goex.IBotInterface
		Pair:      currency.NewPair(currency.BTC, currency.USDT),
		Side:      order.Sell,
		Type:      order.Limit,
		Price:     1000000,
		Amount:    0.1,
		AssetType: asset.Spot,
	}

	// Context will be intercepted when sending an authenticated HTTP request.
	resp, err := b.SubmitOrder(ctx, o)
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println(resp.OrderID)
```
