package ingredients

type Clams interface {
	isClams()
	GetName() string
}

type FreshClams struct {

}

func (clams FreshClams) GetName() string {
	return "Свежие креветки"
}

func (clams FreshClams) isClams()  {

}

type FrozenClams struct {

}

func (f FrozenClams) isClams() {
}

func (f FrozenClams) GetName() string {
	return "Замороженные креветки"
}
