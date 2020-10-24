package ingredients

type Pepperoni interface {
	isPepperoni()
	GetName() string
}

type SlicedPepperoni struct {

}

func (pepperoni SlicedPepperoni) GetName() string {
	return "Sliced pepperoni"
}

func (pepperoni SlicedPepperoni) isPepperoni() {

}
