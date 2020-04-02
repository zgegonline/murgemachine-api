package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/zgegonline/murgemachine-api/model"
)

type Config struct {
	DefaultLight model.Light     `json:"light"`
	Drinks       model.Drinks    `json:"drinks"`
	Cocktails    model.Cocktails `json:"cocktails"`
	Pumps        model.Pumps     `json:"pumps"`
}

func (c Config) getDefaultLight() model.Light {
	return c.DefaultLight
}

func (c Config) getDrinks() model.Drinks {
	return c.Drinks
}
func (c Config) getDrinkList() []model.Drink {
	return c.Drinks.Drinks
}

func (c Config) getCocktails() model.Cocktails {
	return c.Cocktails
}
func (c Config) getCocktailList() []model.Cocktail {
	return c.Cocktails.Cocktails
}
func (c Config) getAvailableCocktails() model.Cocktails {
	return c.Cocktails.GetAvailableCocktails(c.getPumpList())
}
func (c *Config) addCocktail(cocktail model.Cocktail) error {
	if len(cocktail.Ingredients) == 0 {
		return errors.New("Cocktail has no ingredients")
	}
	for _, i := range cocktail.Ingredients {
		ingredientDrink := i.Id
		validCocktail := false
		for _, d := range c.getDrinkList() {
			if ingredientDrink == d.Id {
				validCocktail = true
			}
		}
		if !validCocktail {
			return errors.New("Drink : " + ingredientDrink + " is not present in the config")
		}
	}
	fmt.Println("cocktail " + cocktail.Name + " added")
	c.Cocktails.Cocktails = append(c.getCocktailList(), cocktail)
	return nil
}

func (c Config) getPumps() model.Pumps {
	return c.Pumps
}
func (c Config) getPumpList() []model.Pump {
	return c.Pumps.Pumps
}

func loadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &CurrentConfig.Drinks)
	json.Unmarshal(byteValue, &CurrentConfig.Cocktails)
	json.Unmarshal(byteValue, &CurrentConfig.Pumps)

	jsonFile.Close()

	fmt.Println("Number of drinks loaded : " + strconv.Itoa(len(CurrentConfig.getDrinkList())))
	fmt.Println("Number of cocktails loaded : " + strconv.Itoa(len(CurrentConfig.getCocktailList())))
	fmt.Println("Number of pumps loaded : " + strconv.Itoa(len(CurrentConfig.getPumpList())))
}
