package model

type MqttMessage struct {
	Preparation Preparation `json:"preparation"`
	Light       Light       `json:"light"`
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
