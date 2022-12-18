package apiobjects

import (
	"github.com/vv-projects/go-apicaller/pkg/models/baseobjects"
	"time"
)

type CoinSettings struct {
	General     GeneralSettings
	Name        string
	Description string
	Url         string
	APIKey      string
	Cryptos     []string
}

type CoinResponse struct {
	AssetId            string    `json:"asset_id"`
	Name               string    `json:"name"`
	TypeIsCrypto       int       `json:"type_is_crypto"`
	DataStart          string    `json:"data_start"`
	DataEnd            string    `json:"data_end"`
	DataQuoteStart     time.Time `json:"data_quote_start"`
	DataQuoteEnd       time.Time `json:"data_quote_end"`
	DataOrderbookStart time.Time `json:"data_orderbook_start"`
	DataOrderbookEnd   time.Time `json:"data_orderbook_end"`
	DataTradeStart     time.Time `json:"data_trade_start"`
	DataTradeEnd       time.Time `json:"data_trade_end"`
	DataSymbolsCount   int       `json:"data_symbols_count"`
	Volume1HrsUsd      float64   `json:"volume_1hrs_usd"`
	Volume1DayUsd      float64   `json:"volume_1day_usd"`
	Volume1MthUsd      float64   `json:"volume_1mth_usd"`
	PriceUsd           float64   `json:"price_usd"`
}

type CoinApiResponses []CoinResponse

func (cr CoinApiResponses) ToBase(name string, apiDate string) baseobjects.CryptoObject {
	cryptoMap := make(map[string]float64, len(cr))
	date := ""
	for _, val := range cr {
		cryptoMap[val.Name] = val.PriceUsd
		date = val.DataEnd
	}

	return baseobjects.CryptoObject{
		ApiName:     name,
		DateCall:    apiDate,
		LastUpdated: date,
		CryptoPrice: cryptoMap,
	}
}
