package structs

import (
	"fmt"
	"patterns/observer/interfaces"
)

type ThirdDisplay struct {
	Name string
	Velocity float64
}

func (obj *ThirdDisplay) Update(subject interfaces.Subject) {
	fmt.Println("update third display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.Velocity = weatherData.Value3 * 3
	}
}
func (obj ThirdDisplay) Display() string {
	return fmt.Sprintf("Velocity %6.2f", obj.Velocity)
}