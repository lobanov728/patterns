package beverage

type Mocha struct {
	CostValue float64
	Description string
}
func NewMocha(obj Bevarage) Bevarage {
	mocha := new(Mocha)
	mocha.CostValue += obj.Cost() + 0.20
	mocha.Description = obj.GetDescription() + ", молоко"
	return mocha
}
func (obj Mocha) Cost() float64 {
	return obj.CostValue
}
func (obj Mocha) GetDescription() string {
	return obj.Description
}