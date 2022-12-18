package apiobjects

import (
	"fmt"
	"log"
)

type GeneralSettings struct {
	IsProd         bool
	PreviousPeriod int
}

func FillGeneralSettings(settings map[string]interface{}) GeneralSettings {
	// fill by default
	general := GeneralSettings{
		IsProd:         false,
		PreviousPeriod: 1,
	}
	if v, ok := settings["General Settings"]; ok {
		set := v.(map[string]interface{})
		for key, val := range set {
			switch key {
			case "isprod":
				general.IsProd = val.(bool)
				break
			case "previousperiod":
				pt := val
				fmt.Println(pt)
				general.PreviousPeriod = int(val.(float64))
				break

			default:
				break
			}
		}
	} else {
		log.Println("get general settings by default")
	}
	return general
}
