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

	router := mux.NewRouter()
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/drinks", getDrinks).Methods("GET")
	setup()
	// handleRequests(router)

	fmt.Println("Starting router...")
	return http.ListenAndServe(":8080", router)
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

func handleRequests(router mux.Router) {

}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Home")
}

func getDrinks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Drinks)
}
