package model

type Pumps struct {
	Pumps []Pump `json:"pumps"`
}

type Pump struct {
	Number  int    `json:"number"`
	Content string `json:"content"`
}
