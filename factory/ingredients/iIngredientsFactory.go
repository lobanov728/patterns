package ingredients

type PizzaIngredientsFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateVeggies() []Veggies
	CreatePepperoni() Pepperoni
	CreateClams() Clams
}