package model

type Cocktails struct {
	Cocktails []Cocktail `json:"cocktails"`
}

type Cocktail struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Color       string       `json:"color"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id   string `json:"id"`
	Part int    `json:"part"`
}
