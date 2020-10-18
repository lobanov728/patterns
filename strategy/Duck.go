package strategy

type Duck interface {
	PerformFly()
}

type MallarDuck struct {
	flyBehavior flyBehavior
}
func newMallarDuck() *MallarDuck {
	duck := new(MallarDuck)
	duck.flyBehavior = FlyWithWings{}
	return duck
}
func (o MallarDuck) PerformFly() {
	o.flyBehavior.fly()
}
