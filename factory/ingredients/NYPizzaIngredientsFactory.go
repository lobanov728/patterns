package ingredients

type NYPizzaIngredientsFactory struct {

}

func (Factory *NYPizzaIngredientsFactory) CreateDough() Dough {
	return new(ThinCrustDough)
}
func (Factory *NYPizzaIngredientsFactory) CreateSauce() Sauce {
	return new(MartinaraSauce)
}
func (Factory *NYPizzaIngredientsFactory) CreateCheese() Cheese {
	return new(RegiannoCheese)
}
func (Factory *NYPizzaIngredientsFactory) CreateVeggies() []Veggies {
	return []Veggies{new(Garlic), new(Onion), new(Mushroom), new(RegPepper)}
}
func (Factory *NYPizzaIngredientsFactory) CreatePepperoni() Pepperoni {
	return new(SlicedPepperoni)
}
func (Factory *NYPizzaIngredientsFactory) CreateClams() Clams {
	return new(FreshClams)
}