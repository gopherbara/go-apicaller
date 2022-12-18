package utils

import "github.com/montanaflynn/stats"

func RoundFloat(num float64, params ...int) float64 {
	precision := 2
	if params != nil && params[0] > 0 {
		precision = params[0]
	}
	roundedNum, _ := stats.Round(num, precision)
	return roundedNum
}
