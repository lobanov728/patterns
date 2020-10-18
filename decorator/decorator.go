package main

import "fmt"

type Bevarage interface  {
	cost() float64
	getDescription() string
}
/*************************/
type DarkRoast struct {
	costValue float64
	description string
}
func newDarkRoast() *DarkRoast {
	new := new(DarkRoast)
	new.costValue = 5.99
	new.description = "Кофе глубокой прожарки"
	return new
}
func (obj DarkRoast) cost() float64 {
	return obj.costValue
}
func (obj DarkRoast) getDescription() string {
	return obj.description
}
/*************************/
type Mocha struct {
	costValue float64
	description string
}
func newMocha(bevarage Bevarage) *Mocha {
	new := new(Mocha)
	new.costValue += bevarage.cost() + 0.20
	new.description = bevarage.getDescription() + ", молоко"
	return new
}
func (obj Mocha) cost() float64 {
	return obj.costValue
}
func (obj Mocha) getDescription() string {
	return obj.description
}


func main()  {
	fmt.Println("Decorator")

	darkRoast := newDarkRoast()
	fmt.Println(darkRoast.getDescription())
	fmt.Println(darkRoast.cost())
	mocha := newMocha(darkRoast)
	fmt.Println(mocha.getDescription())
	fmt.Println(mocha.cost())

}
