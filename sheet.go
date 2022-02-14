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
