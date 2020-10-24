package ingredients

type Dough interface {
	isDough()
	GetName() string
}

type ThinCrustDough struct {

}

func (dough ThinCrustDough) GetName() string {
	return "Тонкое хрустящая основа"
}

func (dough ThinCrustDough) isDough() {

}

type ThickDough struct {

}

func (dough ThickDough) GetName() string {
	return "Толстая основа"
}

func (dough ThickDough) isDough() {

}