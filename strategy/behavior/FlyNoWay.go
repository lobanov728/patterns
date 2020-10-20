package behavior

import "fmt"

type FlyNoWay struct {}

func (o FlyNoWay) Fly() {
	fmt.Println("I cannot fly")
}