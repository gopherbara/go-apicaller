package baseobjects

type CryptoObject struct {
	ApiName     string             `json:"apiname,omitempty"`
	DateCall    string             `json:"datecall"`
	LastUpdated string             `json:"lastupdated"`
	CryptoPrice map[string]float64 `json:"cryptoprice,omitempty"`
}
