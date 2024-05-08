package goex

import "strings"

const (
	ExchangeNameAlphapoint  = "alphapoint"
	ExchangeNameBinance     = "binance"
	ExchangeNameBinanceus   = "binanceus"
	ExchangeNameBitfinex    = "bitfinex"
	ExchangeNameBithumb     = "bithumb"
	ExchangeNameBitflyer    = "bitflyer"
	ExchangeNameBitmex      = "bitmex"
	ExchangeNameBitstamp    = "bitstamp"
	ExchangeNameBittrex     = "bittrex"
	ExchangeNameBtcMarkets  = "btc markets"
	ExchangeNameBtse        = "btse"
	ExchangeNameBybit       = "bybit"
	ExchangeNameCoinbasepro = "coinbasepro"
	ExchangeNameCoinut      = "coinut"
	ExchangeNameExmo        = "exmo"
	ExchangeNameGateio      = "gateio"
	ExchangeNameGemini      = "gemini"
	ExchangeNameHitbtc      = "hitbtc"
	ExchangeNameHuobi       = "huobi"
	ExchangeNameItbit       = "itbit"
	ExchangeNameKraken      = "kraken"
	ExchangeNameKucoin      = "kucoin"
	ExchangeNameLbank       = "lbank"
	ExchangeNameOkcoin      = "okcoin"
	ExchangeNameOkx         = "okx"
	ExchangeNamePoloniex    = "poloniex"
	ExchangeNameYobit       = "yobit"
	ExchangeNameZb          = "zb"
)

// IsSupported returns whether or not a specific exchange is supported
func IsSupported(exchangeName string) bool {
	for x := range Exchanges {
		if strings.EqualFold(exchangeName, Exchanges[x]) {
			return true
		}
	}
	return false
}

// Exchanges stores a list of supported exchanges
var Exchanges = []string{
	ExchangeNameAlphapoint,
	ExchangeNameBinance,
	ExchangeNameBinanceus,
	ExchangeNameBitfinex,
	ExchangeNameBithumb,
	ExchangeNameBitflyer,
	ExchangeNameBitmex,
	ExchangeNameBitstamp,
	ExchangeNameBittrex,
	ExchangeNameBtcMarkets,
	ExchangeNameBtse,
	ExchangeNameBybit,
	ExchangeNameCoinbasepro,
	ExchangeNameCoinut,
	ExchangeNameExmo,
	ExchangeNameGateio,
	ExchangeNameGemini,
	ExchangeNameHitbtc,
	ExchangeNameHuobi,
	ExchangeNameItbit,
	ExchangeNameKraken,
	ExchangeNameKucoin,
	ExchangeNameLbank,
	ExchangeNameOkcoin,
	ExchangeNameOkx,
	ExchangeNamePoloniex,
	ExchangeNameYobit,
	ExchangeNameZb,
}
