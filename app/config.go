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
func (c *Config) addDrink(drink model.Drink) error {
	for _, iterDrink := range c.getDrinkList() {
		if iterDrink.Id == drink.Id {
			return errors.New("Drink with id : " + drink.Id + " already exist in the configuration")
		}
		if iterDrink.Name == drink.Name {
			return errors.New("Drink with name : " + drink.Name + " already exist in the configuration")
		}
	}
	return nil
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
	//check name
	for _, iterCocktail := range c.getCocktailList() {
		if iterCocktail.Id == cocktail.Id {
			return errors.New("Cocktail with id : " + strconv.Itoa(cocktail.Id) + " already exist in the configuration")
		}
		if iterCocktail.Name == cocktail.Name {
			return errors.New("Cocktail with name : " + cocktail.Name + " already exist in the configuration")
		}
	}

	// check ingredients
	if len(cocktail.Ingredients) == 0 {
		return errors.New("Cocktail has no ingredients")
	}
	partSum := 0
	for _, i := range cocktail.Ingredients {
		ingredientDrink := i.Id
		partSum += i.Part
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
	if partSum != 100 {
		return errors.New("Sum of the ingredients is not 100 : " + strconv.Itoa(partSum))
	}

	fmt.Println("Adding cocktail " + cocktail.Name + "...")
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
