package fishery

type AddResponse struct {
	UpdatedRange string `json:"updatedRange"`
}

type UpdateResponse struct {
	TotalUpdatedRows int64 `json:"totalUpdatedRows"`
}
