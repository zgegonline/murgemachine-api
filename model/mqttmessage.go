package model

import "fmt"

type MqttMessage struct {
	Preparation Preparation `json:"preparation"`
	Light       Light       `json:"light"`
}

func (msg *MqttMessage) Generate(req Request, cocktails Cocktails, pumps []Pump) error {

	fmt.Println(req.Size)
	msg.Preparation.Size = req.Size
	cocktail, err := cocktails.GetCocktail(req.CocktailId)
	if err != nil {
		return err
	}
	msg.Preparation.PumpsActivation = GetPumpsToActivate(cocktail, pumps)
	if req.Light.Color != "" {
		msg.Light = req.Light
		if msg.Light.Effect == "" {
			msg.Light.Effect = "fixed"
		}
	} else {
		msg.Light.Color = cocktail.Color
		msg.Light.Effect = "fixed"
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
