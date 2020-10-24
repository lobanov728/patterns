package ingredients

type Cheese interface {
	isCheese()
	GetName() string
}

type RegiannoCheese struct {
}

func (cheese RegiannoCheese) GetName() string {
	return "Режианский сыр"
}

func (cheese RegiannoCheese) isCheese() {
}

type MozzarellaCheese struct {
}

func (cheese MozzarellaCheese) GetName() string {
	return "Сыр моцарелла"
}

func (cheese MozzarellaCheese) isCheese() {
}