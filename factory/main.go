package main

import (
	"fmt"
	"patterns/factory/pizzaStore"
)

func main() {
	store := new(pizzaStore.NYPizzaStore)
	pizza := store.CreatePizza("cheese")
	fmt.Println(pizza.GetName())
}
