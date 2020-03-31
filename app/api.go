package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/zgegonline/murgemachine-api/model"

	"github.com/gorilla/mux"
)

var Drinks model.Drinks
var Cocktails model.Cocktails
var Pumps model.Pumps
var MqttClient mqtt.Client

func Start() {
	setup()
	log.Fatal(handleRequests())
}

func setup() {
	loadConfig()
	MqttClient := connectMQTT("192.168.1.100:1883", "murgemachine-api")
	MqttClient.Publish("test", 0, false, "SALUT")
}

func handleRequests() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/drinks", getDrinks).Methods("GET")
	router.HandleFunc("/drink", createDrink).Methods("POST")
	router.HandleFunc("/cocktails", getCocktails).Methods("GET")
	router.HandleFunc("/cocktail", createCocktail).Methods("POST")
	router.HandleFunc("/pumps", getPumps).Methods("GET")
	router.HandleFunc("/mqttmessage", sendMqttMessage).Methods("POST")

	fmt.Println("Starting router...")
	return http.ListenAndServe(":2636", router)
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func getDrinks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Drinks)
}

func getCocktails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Cocktails)
}

func getPumps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pumps)
}

func createDrink(w http.ResponseWriter, r *http.Request) {
	var newDrink model.Drink
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error creating a drink")
	}

	json.Unmarshal(reqBody, &newDrink)

	Drinks.Drinks = append(Drinks.Drinks, newDrink)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDrink)
}

func createCocktail(w http.ResponseWriter, r *http.Request) {
	var newCocktail model.Cocktail
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error creating a cocktail")
	}

	json.Unmarshal(reqBody, &newCocktail)

	Cocktails.Cocktails = append(Cocktails.Cocktails, newCocktail)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCocktail)
}

func sendMqttMessage(w http.ResponseWriter, r *http.Request) {
	var newMqttMessage model.MqttMessage
	var request model.Request

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error sendMqttMessage")
	}

	json.Unmarshal(reqBody, &request)

	newMqttMessage.Preparation.Size = request.Size
	newMqttMessage.Preparation.PumpsActivation = getPumpsToActivate(request.CocktailId)
	newMqttMessage.Light = request.Light

	w.WriteHeader(http.StatusAccepted)

	jsonData, _ := json.Marshal(newMqttMessage)
	MqttClient := connectMQTT("192.168.1.100:1883", "murgemachine-api")
	MqttClient.Publish("murgemachine", 0, false, string(jsonData))
	MqttClient.Disconnect(1000)

	json.NewEncoder(w).Encode(newMqttMessage)
}

func getPumpsToActivate(cocktailId int) []model.PumpActivation {
	cocktail, _ := getCocktail(cocktailId)
	var pumpsToActivate []model.PumpActivation

	for i := 0; i < len(cocktail.Ingredients); i++ {
		ingredient := cocktail.Ingredients[i]

		inPumps := false
		for j := 0; j < len(Pumps.Pumps); j++ {
			if Pumps.Pumps[j].Content == ingredient.Id {
				inPumps = true
				pumpsToActivate = append(pumpsToActivate, model.PumpActivation{Number: Pumps.Pumps[j].Number, Part: ingredient.Part})
			}
		}
		if !inPumps {
			log.Fatal("Can't find drink in pumps")
		}
	}

	return pumpsToActivate
}

func getCocktail(cocktailId int) (model.Cocktail, error) {
	for i := 0; i < len(Cocktails.Cocktails); i++ {
		if cocktailId == Cocktails.Cocktails[i].Id {
			return Cocktails.Cocktails[i], nil
		}
	}
	return model.Cocktail{}, errors.New("Can't find cocktail")
}
