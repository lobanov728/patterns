package pizza

import (
	"fmt"
	"patterns/factory/ingredients"
)

type ClamPizza struct {
	IngredientsFactory ingredients.PizzaIngredientsFactory
	basePizza BasePizza
}

func NewClamPizza(IngredientsFactory ingredients.PizzaIngredientsFactory) Pizza {
	pizza := new(ClamPizza)
	pizza.basePizza = *new(BasePizza)
	pizza.IngredientsFactory = IngredientsFactory

	return pizza
}

func (c *ClamPizza) Prepare() string {
	c.basePizza.dough = c.IngredientsFactory.CreateDough()
	c.basePizza.cheese = c.IngredientsFactory.CreateCheese()
	c.basePizza.sauce = c.IngredientsFactory.CreateSauce()
	c.basePizza.clams = c.IngredientsFactory.CreateClams()
	return fmt.Sprintf(
		"%s. %s",
		c.GetName(),
		c.ToString(),
	)
}

func (c ClamPizza) Bake() string {
	return c.basePizza.Bake()
}

func (c ClamPizza) Cut() string {
	return c.basePizza.Cut()
}

func (c ClamPizza) Box() string {
	return c.basePizza.Box()
}

func (c *ClamPizza) SetName(name string) {
	c.basePizza.SetName(name)
}

func (c *ClamPizza) GetName() string {
	return c.basePizza.GetName()
}

func (c ClamPizza) ToString() string {
	return fmt.Sprintf("%s, %s, %s, %s",
		ingredients.GetIngredientName(c.basePizza.dough),
		ingredients.GetIngredientName(c.basePizza.cheese),
		ingredients.GetIngredientName(c.basePizza.sauce),
		ingredients.GetIngredientName(c.basePizza.clams),
	)
}
