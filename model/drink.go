package model

type Drinks struct {
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
