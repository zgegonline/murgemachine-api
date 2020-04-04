package model

import (
	"errors"
	"fmt"
	"strconv"
)

type Cocktails struct {
	Cocktails []Cocktail `json:"cocktails"`
}

type Cocktail struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Light       Light        `json:"light"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Id   string `json:"id"`
	Part int    `json:"part"`
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

func (c Cocktails) getHighestId() int {
	max := 0
	for _, cocktail := range c.Cocktails {
		if cocktail.Id > max {
			max = cocktail.Id
		}
	}
	return max
}

func (c Cocktails) CheckAndGenerateId(cocktailId int) int {
	idTaken := false
	i := 0
	for ; i < len(c.Cocktails); i++ {
		if c.Cocktails[i].Id == cocktailId {
			idTaken = true
		}
	}
	if idTaken {
		return c.Cocktails[i-1].Id + 1
	} else {
		return cocktailId
	}
}

func (c Cocktails) CheckAndGenerateId2(cocktailId int) int {
	firstAvailableId := 99999
	i := 0
	for ; i < c.getHighestId(); i++ {
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
		return i + 1
	}
}

func (c Cocktails) CheckAndGenerateId3(cocktailId int) int {
	cocktailIdIsTaken := false
	firstAvailableId := -1
	i := 0
	for _, cocktail := range c.Cocktails {
		fmt.Println("i" + strconv.Itoa(i) + "; cocktail.Id" + strconv.Itoa(cocktail.Id))
		if cocktail.Id == cocktailId { //id wanted by user is taken
			cocktailIdIsTaken = true
		} else if !cocktailIdIsTaken && cocktail.Id > cocktailId { //id wanted by user is available
			return cocktailId
		} else if i != cocktail.Id { // i is available for id
			if firstAvailableId == -1 { // the first available id has not been found yet
				firstAvailableId = i
			}
			i = cocktail.Id
		}
		i++
	}
	if !cocktailIdIsTaken {
		return cocktailId
	} else if firstAvailableId == -1 {
		return i
	} else {
		return firstAvailableId
	}
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
