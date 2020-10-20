package beverage

type DarkRoast struct {
	costValue float64
	description string
}
func NewDarkRoast() *DarkRoast {
	darkRoast := new(DarkRoast)
	darkRoast.costValue = 5.99
	darkRoast.description = "Кофе глубокой прожарки"
	return darkRoast
}
func (obj DarkRoast) Cost() float64 {
	return obj.costValue
}
func (obj DarkRoast) GetDescription() string {
	return obj.description
}