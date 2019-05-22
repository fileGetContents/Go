package validations

type Error struct {
	Key string `json:"key"`
	Msg string `json:"msg"`
}

type Errors struct {
	Code int     `json:"code"`
	Es   []Error `json:"es"`
}

func DefaultCode() int {
	return 2000
}
