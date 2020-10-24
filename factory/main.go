package main

import (
	"patterns/factory/ingredients"
	"patterns/factory/pizzaStore"
	"patterns/factory/recipes"
)

func main() {
	nyStore := new(pizzaStore.NYPizzaStore)
	nyStore.CreatePizza("cheese")
	nyStore.CreatePizza("clam")

	chicagoStore := new(pizzaStore.ChicagoStore)
	var r recipes.Recipe
	var emptyRecipe recipes.Recipe
	r = &recipes.VeggiesRecipe{}
	r.SetIngredient(new(ingredients.RegPepper))
	r.SetIngredient(new(ingredients.Onion))
	chicagoStore.CreatePizza("cheese", emptyRecipe)
	chicagoStore.CreatePizza("clam", emptyRecipe)
	chicagoStore.CreatePizza("veggies", r)
}
