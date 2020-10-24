package ingredients

type Clams interface {
	isClams()
	GetName() string
}

type FreshClams struct {

}

func (clams FreshClams) GetName() string {
	return "Fresh clams"
}

func (clams FreshClams) isClams()  {

}