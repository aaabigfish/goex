package engine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/aaabigfish/goex"
	"github.com/aaabigfish/goex/alphapoint"
	"github.com/aaabigfish/goex/binance"
	"github.com/aaabigfish/goex/binanceus"
	"github.com/aaabigfish/goex/bitfinex"
	"github.com/aaabigfish/goex/bitflyer"
	"github.com/aaabigfish/goex/bithumb"
	"github.com/aaabigfish/goex/bitmex"
	"github.com/aaabigfish/goex/bitstamp"
	"github.com/aaabigfish/goex/bittrex"
	"github.com/aaabigfish/goex/btcmarkets"
	"github.com/aaabigfish/goex/btse"
	"github.com/aaabigfish/goex/bybit"
	"github.com/aaabigfish/goex/coinbasepro"
	"github.com/aaabigfish/goex/coinut"
	"github.com/aaabigfish/goex/exmo"
	"github.com/aaabigfish/goex/gateio"
	"github.com/aaabigfish/goex/gemini"
	"github.com/aaabigfish/goex/hitbtc"
	"github.com/aaabigfish/goex/huobi"
	"github.com/aaabigfish/goex/itbit"
	"github.com/aaabigfish/goex/kraken"
	"github.com/aaabigfish/goex/kucoin"
	"github.com/aaabigfish/goex/lbank"
	"github.com/aaabigfish/goex/okcoin"
	"github.com/aaabigfish/goex/okx"
	"github.com/aaabigfish/goex/poloniex"
	"github.com/aaabigfish/goex/yobit"
	"github.com/aaabigfish/goex/zb"
)

var (
	ErrExchangeNotFound = errors.New("exchange not found")
)

// NewExchange helps create a new exchange to be loaded that is
// supported by GCT. This function will return an error if the exchange is not
// supported.
func NewExchange(name string) (goex.Exchange, error) {
	switch strings.ToLower(name) {
	case goex.ExchangeNameAlphapoint:
		return new(alphapoint.Alphapoint), nil
	case goex.ExchangeNameBinanceus:
		return new(binanceus.Binanceus), nil
	case goex.ExchangeNameBinance:
		return new(binance.Binance), nil
	case goex.ExchangeNameBitfinex:
		return new(bitfinex.Bitfinex), nil
	case goex.ExchangeNameBitflyer:
		return new(bitflyer.Bitflyer), nil
	case goex.ExchangeNameBithumb:
		return new(bithumb.Bithumb), nil
	case goex.ExchangeNameBitmex:
		return new(bitmex.Bitmex), nil
	case goex.ExchangeNameBitstamp:
		return new(bitstamp.Bitstamp), nil
	case goex.ExchangeNameBittrex:
		return new(bittrex.Bittrex), nil
	case goex.ExchangeNameBtcMarkets:
		return new(btcmarkets.BTCMarkets), nil
	case goex.ExchangeNameBtse:
		return new(btse.BTSE), nil
	case goex.ExchangeNameBybit:
		return new(bybit.Bybit), nil
	case goex.ExchangeNameCoinut:
		return new(coinut.COINUT), nil
	case goex.ExchangeNameExmo:
		return new(exmo.EXMO), nil
	case goex.ExchangeNameCoinbasepro:
		return new(coinbasepro.CoinbasePro), nil
	case goex.ExchangeNameGateio:
		return new(gateio.Gateio), nil
	case goex.ExchangeNameGemini:
		return new(gemini.Gemini), nil
	case goex.ExchangeNameHitbtc:
		return new(hitbtc.HitBTC), nil
	case goex.ExchangeNameHuobi:
		return new(huobi.HUOBI), nil
	case goex.ExchangeNameItbit:
		return new(itbit.ItBit), nil
	case goex.ExchangeNameKraken:
		return new(kraken.Kraken), nil
	case goex.ExchangeNameKucoin:
		return new(kucoin.Kucoin), nil
	case goex.ExchangeNameLbank:
		return new(lbank.Lbank), nil
	case goex.ExchangeNameOkcoin:
		return new(okcoin.Okcoin), nil
	case goex.ExchangeNameOkx:
		return new(okx.Okx), nil
	case goex.ExchangeNamePoloniex:
		return new(poloniex.Poloniex), nil
	case goex.ExchangeNameYobit:
		return new(yobit.Yobit), nil
	case goex.ExchangeNameZb:
		return new(zb.ZB), nil
	default:
		return nil, fmt.Errorf("'%s', %w", name, ErrExchangeNotFound)
	}
}

func NewExchangeByName(ctx context.Context, name string) (goex.Exchange, error) {
	exch, err := NewExchange(name)
	if err != nil {
		return nil, err
	}

	exch.SetDefaults()

	return exch, nil
}

// NewDefaultExchangeByName returns a defaulted exchange by its name if it
// exists. This will allocate a new exchange and setup the default config for it.
// This will automatically fetch available pairs.
func NewDefaultExchangeByName(ctx context.Context, name string) (goex.Exchange, error) {
	exch, err := NewExchange(name)
	if err != nil {
		return nil, err
	}
	defaultConfig, err := exch.GetDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	err = exch.Setup(defaultConfig)
	if err != nil {
		return nil, err
	}
	return exch, nil
}
