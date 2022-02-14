package fishery

import (
	"fmt"

	"github.com/google/uuid"
)

type Sheet struct {
	client *Client
	name   string
}

const (
	List       = "list"
	OptionArea = "option_area"
	OptionSize = "option_size"
)

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
	for key, _ := range *records {
		(*records)[key].ID = uuid.NewString()
	}

	err = s.client.add(s.name, records, &response)

	if err != nil {
		return
	}

	return
}
