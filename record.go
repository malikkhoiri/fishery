package fishery

import "encoding/json"

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
