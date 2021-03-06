package model

import "fmt"

type MqttMessage struct {
	Preparation Preparation `json:"preparation"`
	Light       Light       `json:"light"`
}

func (msg *MqttMessage) Generate(req Request, cocktails Cocktails, pumps Pumps) error {

	fmt.Println(req.Size)
	msg.Preparation.Size = req.Size
	cocktail, err := cocktails.GetCocktail(req.CocktailId)
	if err != nil {
		return err
	}
	msg.Preparation.PumpsActivation, err = pumps.GetPumpsToActivate(cocktail)
	if err != nil {
		return err
	}
	if req.Light.Color != "" {
		msg.Light = req.Light
		if msg.Light.Effect == "" {
			msg.Light.Effect = "fixed"
		}
	} else {
		msg.Light = cocktail.Light
	}
	return nil
}

type Preparation struct {
	Size            int              `json:"size"`
	PumpsActivation []PumpActivation `json:"pumpsActivation"`
}

type PumpActivation struct {
	Number int `json:"number"`
	Part   int `json:"part"`
}

type Light struct {
	Color  string `json:"color"`
	Effect string `json:"effect"`
}

type Request struct {
	CocktailId int   `json:"cocktailId"`
	Size       int   `json:"size"`
	Light      Light `json:"light"`
}
