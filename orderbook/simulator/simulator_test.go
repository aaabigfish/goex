package simulator

import (
	"context"
	"testing"

	"github.com/aaabigfish/goex/asset"
	"github.com/aaabigfish/goex/bitstamp"
	"github.com/aaabigfish/goex/common/convert"
	"github.com/aaabigfish/goex/currency"
)

func TestSimulate(t *testing.T) {
	b := bitstamp.Bitstamp{}
	b.SetDefaults()
	b.Verbose = false
	b.CurrencyPairs = currency.PairsManager{
		UseGlobalFormat: true,
		RequestFormat: &currency.PairFormat{
			Uppercase: true,
		},
		Pairs: map[asset.Item]*currency.PairStore{
			asset.Spot: {
				AssetEnabled: convert.BoolPtr(true),
			},
		},
	}
	o, err := b.FetchOrderbook(context.Background(),
		currency.NewPair(currency.BTC, currency.USD), asset.Spot)
	if err != nil {
		t.Fatal(err)
	}
	_, err = o.SimulateOrder(10000000, true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = o.SimulateOrder(2171, false)
	if err != nil {
		t.Fatal(err)
	}
}
