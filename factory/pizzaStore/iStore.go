package pizzaStore

import (
	"fmt"
	"patterns/factory/ingredients"
	"patterns/factory/pizza"
)

type PizzaStore interface {
	CreatePizza(name string) pizza.Pizza
}

type NYPizzaStore struct {
}

func (N NYPizzaStore) CreatePizza(name string) pizza.Pizza {
	NYIngredientsFactory := new(ingredients.NYPizzaIngredientsFactory)

	var pizzaObj pizza.Pizza

	switch name {
	case "cheese":
		pizzaObj = pizza.NewCheesePizza(NYIngredientsFactory)
		pizzaObj.SetName("Нью Йоркская сырная пицца")
		fmt.Println(pizzaObj.Prepare())
		fmt.Println(pizzaObj.Bake())
		fmt.Println(pizzaObj.Box())
		fmt.Println(pizzaObj.Cut())
	}

	return pizzaObj
}

