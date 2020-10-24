package recipes

import (
	"patterns/factory/ingredients"
)

type Recipe interface {
	SetIngredient(obj ingredients.Ingredient)
	GetIngredients() []ingredients.Ingredient
}

type VeggiesRecipe struct {
	IngredientsList []ingredients.Veggies
}

func (c *VeggiesRecipe) SetIngredient(obj ingredients.Ingredient) {
	veggies, ok := obj.(ingredients.Veggies)
	if ok {
		c.IngredientsList = append(c.IngredientsList, veggies)
	}
}

func (c VeggiesRecipe) GetIngredients() []ingredients.Ingredient {
	var result []ingredients.Ingredient
	for _, veggies := range c.IngredientsList {
		result = append(result, veggies.(ingredients.Ingredient))
	}
	return result
}
