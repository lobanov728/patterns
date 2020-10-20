package behavior

import "fmt"

type FlyWithWings struct {}

func (o FlyWithWings) Fly() {
	fmt.Println("Fly by wings")
}
