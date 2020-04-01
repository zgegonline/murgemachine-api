package model

import (
	"errors"
	"fmt"
	"strconv"
)

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

func (c Cocktails) GetAvailableCocktails(pumps []Pump) Cocktails {
	var availableCocktails []Cocktail
	for i := 0; i < len(c.Cocktails); i++ {
		currentCocktail := c.Cocktails[i]
		if currentCocktail.Available(pumps) {
			availableCocktails = append(availableCocktails, currentCocktail)
		}
	}
	return Cocktails{availableCocktails}
}

func (c Cocktails) CheckAndGenerateId(cocktailId int) int {
	firstAvailableId := 99999
	i := 0
	for ; i < len(c.Cocktails); i++ {
		if i != c.Cocktails[i].Id { // id is available
			if i == cocktailId { // id wanted by user is available
				return cocktailId
			} else if i < firstAvailableId { // we have found the first available id
				firstAvailableId = i
			}
		}
	}
	fmt.Println("i=" + strconv.Itoa(i) + "; cocktailId=" + strconv.Itoa(cocktailId) + "; firstAvailableId=" + strconv.Itoa(firstAvailableId))
	if cocktailId > i {
		return cocktailId
	} else if firstAvailableId != 99999 {
		return firstAvailableId
	} else {
		return i
	}
}

type Cocktail struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Color       string       `json:"color"`
	Ingredients []Ingredient `json:"ingredients"`
}

func (cocktail Cocktail) Available(pumps []Pump) bool {
	for i := 0; i < len(cocktail.Ingredients); i++ {
		drinkId := cocktail.Ingredients[i].Id
		drinkFound := false
		for j := 0; j < len(pumps); j++ {
			if pumps[j].Content == drinkId {
				drinkFound = true
				break
			}
		}
		if !drinkFound {
			return false
		}
	}
	return true
}

type Ingredient struct {
	Id   string `json:"id"`
	Part int    `json:"part"`
}
