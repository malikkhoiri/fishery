package fishery

type Sheet struct {
	client *Client
	name   string
}

const (
	List       = "list"
	OptionArea = "option_area"
	OptionSize = "option_size"
)
