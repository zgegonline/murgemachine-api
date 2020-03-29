package model

type Drinks struct {
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Type DrinkType `json:"type"`
}

type DrinkType int

const (
	Alcohol DrinkType = 0
	Soft    DrinkType = 1
)
