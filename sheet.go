package fishery

import (
	"fmt"

	"github.com/google/uuid"
)

type Sheet struct {
	client *Client
	name   string
}

type FilterArea struct {
	ProvinceArea string `json:"area_provinsi,omitempty"`
	CityArea     string `json:"area_kota,omitempty"`
}

const (
	List       = "list"
	OptionArea = "option_area"
	OptionSize = "option_size"
)

func (s *Sheet) GetAll() (records *Records, err error) {
	err = s.client.getAll(s.name, &records)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) GetByID(id string) (record *Record, err error) {
	search := Search{"uuid": id}
	records := make(Records, 0)
	err = s.client.get(s.name, search, &records)

	if err != nil {
		return
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	record = &records[0]
	return
}

func (s *Sheet) AddRecords(records *Records) (response *AddResponse, err error) {
	for key := range *records {
		(*records)[key].ID = uuid.NewString()
	}

	err = s.client.add(s.name, records, &response)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) UpdateRecords(record *Record) (response *UpdateResponse, err error) {
	reqData := &UpdateRequest{
		Condition: map[string]interface{}{"uuid": record.ID},
		Set:       record,
	}

	err = s.client.update(s.name, reqData, &response)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) DeleteRecords(id string) (response *DeleteResponse, err error) {
	reqData := &DeleteRequest{
		Condition: map[string]interface{}{"uuid": id},
	}

	err = s.client.delete(s.name, reqData, &response)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) GetAllByCommodity(commodity string) (records *Records, err error) {
	search := Search{"komoditas": commodity}
	err = s.client.get(s.name, search, &records)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) GetAllByArea(fa FilterArea) (records *Records, err error) {
	err = s.client.get(s.name, fa, &records)

	if err != nil {
		return
	}

	return
}

func (s *Sheet) GetMaxPriceByWeek(week int) (record *Record, err error) {
	records := &Records{}
	err = s.client.getAll(s.name, records)

	if err != nil {
		return
	}

	return records.getMaxPriceByWeek(week)
}

func (s *Sheet) GetMaxPriceByCommudity(commodity string) (record *Record, err error) {
	records := &Records{}
	err = s.client.getAll(s.name, records)

	if err != nil {
		return
	}

	return records.getMaxPriceByCommodity(commodity)
}

func (s *Sheet) GetMostRecords() (result string, err error) {
	records := &Records{}
	err = s.client.getAll(s.name, records)

	if err != nil {
		return
	}

	return records.getMostRecords(), nil
}

func (s *Sheet) GetLatestRecords() (records *Records, err error) {
	err = s.client.getAll(s.name, &records)

	if err != nil {
		return
	}

	return records.getLatestRecords(), nil
}
