package model

import "errors"

type Cocktails struct {
	Cocktails []Cocktail `json:"cocktails"`
}

func (c Cocktails) GetCocktail(cocktailId int) (Cocktail, error) {
	for i := 0; i < len(c.Cocktails); i++ {
		if cocktailId == c.Cocktails[i].Id {
			return c.Cocktails[i], nil
		}
	}
	return Cocktail{}, errors.New("Can't find cocktail")
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
