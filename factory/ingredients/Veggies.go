package ingredients

type Veggies interface {
	isVeggies()
	GetName() string
}

type Garlic struct {

}

func (obj Garlic) GetName() string {
	return "Garlic"
}

func (obj Garlic) isVeggies() {

}

type Onion struct {

}

func (obj Onion) GetName() string {
	return "Onion"
}

func (obj Onion) isVeggies() {

}

type Mushroom struct {

}

func (obj Mushroom) GetName() string {
	return "Mushroom"
}

func (obj Mushroom) isVeggies() {

}

type RegPepper struct {

}

func (obj RegPepper) GetName() string {
	return "Red pepper"
}

func (obj RegPepper) isVeggies() {

}