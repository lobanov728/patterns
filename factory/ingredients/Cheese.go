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

func NewRegiannoCheese() *RegiannoCheese {
	return new(RegiannoCheese)
}

func (cheese RegiannoCheese) isCheese() {
}
