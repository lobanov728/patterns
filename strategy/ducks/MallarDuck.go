package ducks

import (
	"patterns/strategy/behavior"
)

type MallarDuck struct {
	FlyBehavior behavior.FlyBehavior
}
func NewMallarDuck() *MallarDuck {
	duck := new(MallarDuck)
	duck.FlyBehavior = behavior.FlyWithWings{}
	return duck
}
func (o MallarDuck) PerformFly() {
	o.FlyBehavior.Fly()
}
