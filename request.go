package fishery

type UpdateRequest struct {
	Condition map[string]interface{} `json:"condition"`
	Set       interface{}            `json:"set"`
}
