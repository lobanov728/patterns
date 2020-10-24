package pizzaStore

import (
	"fmt"
	"patterns/factory/ingredients"
	"patterns/factory/pizza"
)

type NYPizzaStore struct {
}
func (N NYPizzaStore) CreatePizza(name string) pizza.Pizza {
	ingredientsFactory := new(ingredients.NYPizzaIngredientsFactory)

	var pizzaObj pizza.Pizza
	switch name {
	case "cheese":
		pizzaObj = pizza.NewCheesePizza(ingredientsFactory)
		pizzaObj.SetName("Нью Йоркская сырная пицца")
	case "clam":
		pizzaObj = pizza.NewClamPizza(ingredientsFactory)
		pizzaObj.SetName("Нью Йоркская пицца с креветками")
	}
	fmt.Println(pizzaObj.Prepare())
	fmt.Println(pizzaObj.Bake())
	fmt.Println(pizzaObj.Box())
	fmt.Println(pizzaObj.Cut())


	return pizzaObj
}
