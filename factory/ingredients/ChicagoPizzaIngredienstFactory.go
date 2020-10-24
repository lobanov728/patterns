package ingredients

type ChicagoPizzaIngredienstFactory struct {

}

func (Factory *ChicagoPizzaIngredienstFactory) CreateDough() Dough {
	return new(ThickDough)
}
func (Factory *ChicagoPizzaIngredienstFactory) CreateSauce() Sauce {
	return new(PlumTomatoSauce)
}
func (Factory *ChicagoPizzaIngredienstFactory) CreateCheese() Cheese {
	return new(MozzarellaCheese)
}
func (Factory *ChicagoPizzaIngredienstFactory) CreateVeggies() []Veggies {
	return []Veggies{new(Garlic), new(Onion), new(Mushroom), new(RegPepper)}
}
func (Factory *ChicagoPizzaIngredienstFactory) CreatePepperoni() Pepperoni {
	return new(SlicedPepperoni)
}
func (Factory *ChicagoPizzaIngredienstFactory) CreateClams() Clams {
	return new(FrozenClams)
}