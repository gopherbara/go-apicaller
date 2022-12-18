package baseobjects

type StocksObject struct {
	ApiName     string             `json:"apiname,omitempty"`
	DateCall    string             `json:"datecall,omitempty"`
	LastUpdated string             `json:"lastupdated,omitempty"`
	Stock       string             `json:"stock,omitempty"`
	Price       float64            `json:"price,omitempty"`
	StockPrice  map[string]float64 `json:"stockprice,omitempty"`
}
