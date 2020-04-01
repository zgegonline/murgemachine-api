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
	defaultLight model.Light     `json:"light"`
	drinks       model.Drinks    `json:"drinks"`
	cocktails    model.Cocktails `json:"cocktails"`
	pumps        model.Pumps     `json:"pumps"`
}

func (c Config) getDefaultLight() model.Light {
	return c.defaultLight
}

func (c Config) getDrinks() model.Drinks {
	return c.drinks
}
func (c Config) getDrinkList() []model.Drink {
	return c.drinks.Drinks
}

func (c Config) getCocktails() model.Cocktails {
	return c.cocktails
}
func (c Config) getCocktailList() []model.Cocktail {
	return c.cocktails.Cocktails
}
func (c Config) getAvailableCocktails() model.Cocktails {
	return c.cocktails.GetAvailableCocktails(c.getPumpList())
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
	fmt.Println(strconv.Itoa(len(c.getCocktailList())))
	c.cocktails.Cocktails = append(c.getCocktailList(), cocktail)
	fmt.Println(strconv.Itoa(len(c.getCocktailList())))
	return nil
}

func (c Config) getPumps() model.Pumps {
	return c.pumps
}
func (c Config) getPumpList() []model.Pump {
	return c.pumps.Pumps
}

func loadConfig() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &CurrentConfig.drinks)
	json.Unmarshal(byteValue, &CurrentConfig.cocktails)
	json.Unmarshal(byteValue, &CurrentConfig.pumps)

	jsonFile.Close()

	fmt.Println("Number of drinks loaded : " + strconv.Itoa(len(CurrentConfig.getDrinkList())))
	fmt.Println("Number of cocktails loaded : " + strconv.Itoa(len(CurrentConfig.getCocktailList())))
	fmt.Println("Number of pumps loaded : " + strconv.Itoa(len(CurrentConfig.getPumpList())))
}
