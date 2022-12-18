package baseobjects

type CurrencyObject struct {
	ApiName      string             `json:"apiname,omitempty"`
	DateCall     string             `json:"datecall"`
	FromCurrency string             `json:"fromcurrency,omitempty"`
	LastUpdated  string             `json:"lastupdated"`
	Rates        map[string]float64 `json:"rates,omitempty"`
}
