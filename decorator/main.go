package main

import (
	"fmt"
	"patterns/decorator/beverage"
)

func main()  {
	fmt.Println("Decorator")
	darkRoast := beverage.NewDarkRoast()
	fmt.Println(darkRoast.GetDescription())
	fmt.Println(darkRoast.Cost())
	mocha := beverage.NewMocha(darkRoast)
	fmt.Println(mocha.GetDescription())
	fmt.Println(mocha.Cost())
}
