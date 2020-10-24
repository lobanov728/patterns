package pizza

import (
	"fmt"
	"patterns/factory/ingredients"
)

type CheesePizza struct {
	IngredientsFactory ingredients.PizzaIngredientsFactory
	basePizza BasePizza
}

func NewCheesePizza(IngredientsFactory ingredients.PizzaIngredientsFactory) Pizza {
	pizza := new(CheesePizza)
	pizza.basePizza = *new(BasePizza)
	pizza.IngredientsFactory = IngredientsFactory

	return pizza
}

func (c *CheesePizza) Prepare() string {
	c.basePizza.dough = c.IngredientsFactory.CreateDough()
	c.basePizza.cheese = c.IngredientsFactory.CreateCheese()
	c.basePizza.sauce = c.IngredientsFactory.CreateSauce()
	return fmt.Sprintf(
		"%s. %s",
		c.GetName(),
		c.ToString(),
	)
}

func (c CheesePizza) Bake() string {
	return c.basePizza.Bake()
}

func (c CheesePizza) Cut() string {
	return c.basePizza.Cut()
}

func (c CheesePizza) Box() string {
	return c.basePizza.Box()
}

func (c *CheesePizza) SetName(name string) {
	c.basePizza.SetName(name)
}

func (c *CheesePizza) GetName() string {
	return c.basePizza.GetName()
}

func (c CheesePizza) ToString() string {
	return fmt.Sprintf("%s, %s, %s",
		getIngredientName(c.basePizza.dough),
		getIngredientName(c.basePizza.cheese),
		getIngredientName(c.basePizza.sauce),
	);
}

func getIngredientName(obj ingredients.Ingredient) string {
	return obj.GetName()
}
