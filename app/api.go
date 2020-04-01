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

var CurrentConfig Config
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
	json.NewEncoder(w).Encode(CurrentConfig.getDrinks())
}

func getCocktails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CurrentConfig.getCocktails())
}

func getAvailableCocktails(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CurrentConfig.getAvailableCocktails())
}

func getPumps(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(CurrentConfig.getPumps())
}

func createDrink(w http.ResponseWriter, r *http.Request) {
	var newDrink model.Drink
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error creating a drink")
	}

	json.Unmarshal(reqBody, &newDrink)

	CurrentConfig.drinks.Drinks = append(CurrentConfig.getDrinkList(), newDrink)

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

	newCocktail.Id = CurrentConfig.getCocktails().CheckAndGenerateId(newCocktail.Id)

	err = CurrentConfig.addCocktail(newCocktail)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newCocktail)
}

func requestCocktail(w http.ResponseWriter, r *http.Request) {
	var newMqttMessage model.MqttMessage
	var request model.Request

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error sendMqttMessage")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.Unmarshal(reqBody, &request)
	err = newMqttMessage.Generate(request, CurrentConfig.getCocktails(), CurrentConfig.getPumps())
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)

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
		c, _ := CurrentConfig.getCocktails().GetCocktail(cocktailId)
		return model.Light{Color: c.Color, Effect: "fixed"}
	}
}
