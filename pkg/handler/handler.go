package handler

type Handler interface {
	Handle()
	Send()
}

type GrowthObject struct {
	ObjectChange  string  `json:"object_change"`
	PreviousValue float64 `json:"previous_value"`
	TodayValue    float64 `json:"today_value"`
	Growth        float64 `json:"growth"`
}
