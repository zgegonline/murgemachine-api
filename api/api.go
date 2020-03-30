package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/zgegonline/murgemachine-restapi/model"

	"github.com/gorilla/mux"
)

var Drinks model.Drinks
var Cocktails model.Cocktails
var Pumps model.Pumps

func Start() error {
	setup()
	return handleRequests()
}

func setup() {
	loadConfig()
}

func loadConfig() {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &Drinks)
	json.Unmarshal(byteValue, &Cocktails)
	json.Unmarshal(byteValue, &Pumps)

	jsonFile.Close()

	fmt.Println("Number of drinks loaded : " + strconv.Itoa(len(Drinks.Drinks)))
	fmt.Println("Number of cocktails loaded : " + strconv.Itoa(len(Cocktails.Cocktails)))
	fmt.Println("Number of pumps loaded : " + strconv.Itoa(len(Pumps.Pumps)))
}

func handleRequests() error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/drinks", getDrinks).Methods("GET")
	router.HandleFunc("/drink", createDrink).Methods("POST")
	router.HandleFunc("/cocktails", getCocktails).Methods("GET")
	router.HandleFunc("/pumps", getPumps).Methods("GET")

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
		fmt.Println(err)
		fmt.Fprintf(w, "Error creating a drink")
	}

	json.Unmarshal(reqBody, &newDrink)
	fmt.Println(newDrink)

	Drinks.Drinks = append(Drinks.Drinks, newDrink)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDrink)
}
