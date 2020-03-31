package app

import (
	"encoding/json"
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
	router.HandleFunc("/available-cocktails", getAvailableCocktails).Methods("GET")
	router.HandleFunc("/cocktail", createCocktail).Methods("POST")
	router.HandleFunc("/pumps", getPumps).Methods("GET")
	router.HandleFunc("/request-cocktail", requestCocktail).Methods("POST")

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

func getAvailableCocktails(w http.ResponseWriter, r *http.Request) {
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

func requestCocktail(w http.ResponseWriter, r *http.Request) {
	var newMqttMessage model.MqttMessage
	var request model.Request

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error sendMqttMessage")
	}

	json.Unmarshal(reqBody, &request)

	newMqttMessage.Generate(request, Cocktails, Pumps.Pumps)

	w.WriteHeader(http.StatusAccepted)

	jsonData, _ := json.Marshal(newMqttMessage)
	MqttClient := connectMQTT("192.168.1.100:1883", "murgemachine-api")
	MqttClient.Publish("murgemachine", 0, false, string(jsonData))
	MqttClient.Disconnect(1000)

	json.NewEncoder(w).Encode(newMqttMessage)
}

// return light param if it not empty, otherwise return Light {color : Cocktail.Color, effect : "fixed"}
func getLight(cocktailId int, light model.Light) model.Light {
	if light.Color != "" {
		return light
	} else {
		c, _ := Cocktails.GetCocktail(cocktailId)
		return model.Light{Color: c.Color, Effect: "fixed"}
	}
}
