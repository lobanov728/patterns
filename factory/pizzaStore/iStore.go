package pizzaStore

import (
	"patterns/factory/pizza"
)

type PizzaStore interface {
	CreatePizza(name string) pizza.Pizza
}
