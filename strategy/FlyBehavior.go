package strategy

import "fmt"

type flyBehavior interface {
	fly()
}

type FlyWithWings struct {}
func (o FlyWithWings) fly() {
	fmt.Println("Fly by wings")
}

type FlyNoWay struct {}
func (o FlyNoWay) fly() {
	fmt.Println("I cannot fly")
}