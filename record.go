package fishery

import (
	"encoding/json"
	"fmt"
	"time"
)

type Record struct {
	ID           string      `json:"uuid"`
	Commodity    string      `json:"komoditas"`
	ProvinceArea string      `json:"area_provinsi"`
	CityArea     string      `json:"area_kota"`
	Size         json.Number `json:"size"`
	Price        json.Number `json:"price"`
	DateParsed   string      `json:"tgl_parsed"`
	Timestamp    json.Number `json:"timestamp"`
}

type Records []Record

func (rs *Records) getMaxPriceByWeek(week int) (result *Record, err error) {
	var maxPrice int64 = 0
	for key := range *rs {
		t, _ := (*rs)[key].Timestamp.Int64()
		_, w := time.Unix(t, 0).ISOWeek()

		if w > week {
			continue
		}

		p, _ := (*rs)[key].Price.Int64()

		if p > maxPrice {
			maxPrice = p
			result = &(*rs)[key]
		}
	}

	if result == nil {
		return result, fmt.Errorf("no record for %d weeks", week)
	}

	return
}

func (rs *Records) getMaxPriceByCommodity(commodity string) (result *Record, err error) {
	var maxPrice int64 = 0
	var c string
	for key := range *rs {
		c = (*rs)[key].Commodity

		if c != commodity {
			continue
		}

		p, _ := (*rs)[key].Price.Int64()

		if p > maxPrice {
			maxPrice = p
			result = &(*rs)[key]
		}
	}

	if result == nil {
		return result, fmt.Errorf("no record for '%s' commodity", commodity)
	}

	return
}
