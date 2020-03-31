package model

import (
	"log"
)

func Available(cocktail Cocktail, pumps []Pump) bool {
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

func GetPumpsToActivate(cocktail Cocktail, pumps []Pump) []PumpActivation {
	var pumpsToActivate []PumpActivation

	for i := 0; i < len(cocktail.Ingredients); i++ {
		ingredient := cocktail.Ingredients[i]

		inPumps := false
		for j := 0; j < len(pumps); j++ {
			if pumps[j].Content == ingredient.Id {
				inPumps = true
				pumpsToActivate = append(pumpsToActivate, PumpActivation{Number: pumps[j].Number, Part: ingredient.Part})
			}
		}
		if !inPumps {
			log.Fatal("Can't find drink in pumps")
		}
	}

	return pumpsToActivate
}
