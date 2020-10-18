package strategy

func Run()  {
	malarDuck := newMallarDuck()
	malarDuck.PerformFly()
	malarDuck.flyBehavior = FlyNoWay{}

	malarDuck.PerformFly()
}
