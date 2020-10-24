package pizza

import (
	"fmt"
	"patterns/factory/ingredients"
	"patterns/factory/recipes"
	"strings"
)

type VeggiesPizza struct {
	IngredientsFactory ingredients.PizzaIngredientsFactory
	basePizza BasePizza
	recipe recipes.Recipe
}

func NewVeggiesPizza(IngredientsFactory ingredients.PizzaIngredientsFactory, r recipes.Recipe) Pizza {
	pizza := new(VeggiesPizza)
	pizza.basePizza = *new(BasePizza)
	pizza.IngredientsFactory = IngredientsFactory
	pizza.recipe = r
	return pizza
}

func (c *VeggiesPizza) Prepare() string {
	c.basePizza.dough = c.IngredientsFactory.CreateDough()
	c.basePizza.cheese = c.IngredientsFactory.CreateCheese()
	c.basePizza.sauce = c.IngredientsFactory.CreateSauce()
	c.basePizza.clams = c.IngredientsFactory.CreateClams()
	veggies := c.IngredientsFactory.CreateVeggies()
	for _, v := range veggies  {
		for _, v1 := range c.recipe.GetIngredients() {
			if v == v1 {
				c.basePizza.veggies = append(c.basePizza.veggies, v1.(ingredients.Veggies))
			}
		}
	}


	return fmt.Sprintf(
		"%s. %s",
		c.GetName(),
		c.ToString(),
	)
}

func (c VeggiesPizza) Bake() string {
	return c.basePizza.Bake()
}

func (c VeggiesPizza) Cut() string {
	return c.basePizza.Cut()
}

func (c VeggiesPizza) Box() string {
	return c.basePizza.Box()
}

func (c *VeggiesPizza) SetName(name string) {
	c.basePizza.SetName(name)
}

func (c *VeggiesPizza) GetName() string {
	return c.basePizza.GetName()
}

func (c VeggiesPizza) ToString() string {
	var sauceStr []string
	for _, veggies := range c.basePizza.veggies {
		sauceStr = append(sauceStr, ingredients.GetIngredientName(veggies))
	}
	return fmt.Sprintf("%s, %s, %s, %s",
		ingredients.GetIngredientName(c.basePizza.dough),
		ingredients.GetIngredientName(c.basePizza.cheese),
		ingredients.GetIngredientName(c.basePizza.sauce),
		strings.Join(sauceStr, ", "),
	)
}
