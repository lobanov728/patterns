package main

import (
	"patterns/strategy/behavior"
	"patterns/strategy/ducks"
)

func main()  {
	malarDuck := ducks.NewMallarDuck()
	malarDuck.PerformFly()
	malarDuck.FlyBehavior = behavior.FlyNoWay{}
	malarDuck.PerformFly()
}
