package ingredients

type Sauce interface {
	isSauce()
	GetName() string
}

type MartinaraSauce struct {

}

func (sauce MartinaraSauce) GetName() string {
	return "Соус маринара"
}

func (sauce MartinaraSauce) isSauce() {

}

