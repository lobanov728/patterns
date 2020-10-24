package ingredients

type Veggies interface {
	isVeggies()
	GetName() string
}

type Garlic struct {

}

func (obj Garlic) GetName() string {
	return "Чеснок"
}

func (obj Garlic) isVeggies() {

}

type Onion struct {

}

func (obj Onion) GetName() string {
	return "Лук"
}

func (obj Onion) isVeggies() {

}

type Mushroom struct {

}

func (obj Mushroom) GetName() string {
	return "Грибы"
}

func (obj Mushroom) isVeggies() {

}

type RegPepper struct {

}

func (obj RegPepper) GetName() string {
	return "Красный перец"
}

func (obj RegPepper) isVeggies() {

}