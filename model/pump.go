package model

import "errors"

type Pumps struct {
	Pumps []Pump `json:"pumps"`
}

type Pump struct {
	Number  int    `json:"number"`
	Content string `json:"content"`
}

func (pumps Pumps) GetPumpsToActivate(cocktail Cocktail) ([]PumpActivation, error) {
	var pumpsToActivate []PumpActivation

	for i := 0; i < len(cocktail.Ingredients); i++ {
		ingredient := cocktail.Ingredients[i]

		inPumps := false
		for j := 0; j < len(pumps.Pumps); j++ {
			if pumps.Pumps[j].Content == ingredient.Id {
				inPumps = true
				pumpsToActivate = append(pumpsToActivate, PumpActivation{Number: pumps.Pumps[j].Number, Part: ingredient.Part})
			}
		}
		if !inPumps {
			return nil, errors.New("Drink " + ingredient.Id + " is not available in pumps")
		}
	}

	return pumpsToActivate, nil
}
