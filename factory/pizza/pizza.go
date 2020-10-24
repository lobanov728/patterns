package pizza

import (
	"fmt"
	"patterns/factory/ingredients"
)

type Pizza interface {
	Prepare() string
	Bake() string
	Cut() string
	Box() string
	SetName(name string)
	GetName() string
	ToString() string
}

type BasePizza struct {
	name string
	dough ingredients.Dough
	sauce ingredients.Sauce
	cheese ingredients.Cheese
	veggies []ingredients.Veggies
	pepperoni ingredients.Pepperoni
	clams ingredients.Clams
}

func (b *BasePizza) Prepare() string {
	return ""
}

func (b BasePizza) Bake() string {
	return fmt.Sprint("Готовить при 350 градусах 25 минут")
}

func (b BasePizza) Cut() string {
	return fmt.Sprint("Разрезать по диагонали")
}

func (b BasePizza) Box() string {
	return fmt.Sprint("Упаковать в фирменную коробку")
}

func (b *BasePizza) SetName(name string) {
	b.name = name
}

func (b *BasePizza) GetName() string {
	return b.name
}

func (b BasePizza) ToString() string {
	return ""
}
