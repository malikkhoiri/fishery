package fishery

import (
	"encoding/json"
	"fmt"
	"sort"
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

func (rs *Records) getMostRecords() (result string) {
	m := map[string]int{}
	var maxCount int
	var f string
	for _, a := range *rs {
		m[a.Commodity]++
		if m[a.Commodity] > maxCount {
			maxCount = m[a.Commodity]
			f = a.Commodity
		}
	}

	return f
}

func (rs *Records) getRangePrice(from, to int64) (records *Records) {
	newRecords := make(Records, 0)

	for key := range *rs {
		price, _ := (*rs)[key].Price.Int64()
		fmt.Println(price)
		if price < from && price > to {
			continue
		}
		newRecords = append(newRecords, (*rs)[key])
	}

	return &newRecords
}

func (rs Records) getLatestRecords() (records *Records) {
	sort.Sort(Records(rs))
	newRecords := rs[:10]
	return &newRecords
}

func (rs Records) Len() int {
	return len(rs)
}

func (rs Records) Less(i, j int) bool {
	ti, _ := rs[i].Timestamp.Int64()
	rsi := time.Unix(ti, 0)
	tj, _ := rs[j].Timestamp.Int64()
	rsj := time.Unix(tj, 0)
	return rsi.After(rsj)
}

func (rs Records) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}
