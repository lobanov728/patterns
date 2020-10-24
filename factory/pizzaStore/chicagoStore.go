package pizzaStore

import (
	"fmt"
	"patterns/factory/ingredients"
	"patterns/factory/pizza"
	"patterns/factory/recipes"
)

type ChicagoStore struct {
}
func (c ChicagoStore) CreatePizza(name string, r recipes.Recipe) pizza.Pizza {
	ingredientsFactory := new(ingredients.ChicagoPizzaIngredienstFactory)

	var pizzaObj pizza.Pizza
	switch name {
	case "cheese":
		pizzaObj = pizza.NewCheesePizza(ingredientsFactory)
		pizzaObj.SetName("Чикагская сырная пицца")
	case "clam":
		pizzaObj = pizza.NewClamPizza(ingredientsFactory)
		pizzaObj.SetName("Чикагская пицца с креветками")
	case "veggies":
		pizzaObj = pizza.NewVeggiesPizza(ingredientsFactory, r)
		pizzaObj.SetName("Чикагская пицца с овощами")
	}
	fmt.Println(pizzaObj.Prepare())
	fmt.Println(pizzaObj.Bake())
	fmt.Println(pizzaObj.Box())
	fmt.Println(pizzaObj.Cut())


	return pizzaObj
}
