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

type PlumTomatoSauce struct {

}

func (p PlumTomatoSauce) isSauce() {
}

func (p PlumTomatoSauce) GetName() string {
	return "Сливочно томатный соус"
}


