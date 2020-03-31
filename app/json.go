package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func loadConfig() {
	jsonFile, err := os.Open("config.json")
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
