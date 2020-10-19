package structs

import (
	"fmt"
	"patterns/observer/interfaces"
)

type FirstDisplay struct {
	Name              string
	Temperature       float64
}
func (obj *FirstDisplay) Update(subject interfaces.Subject) {
	fmt.Println("update first display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.Temperature = weatherData.Value1
	}
}
func (obj FirstDisplay) Display() string {
	return fmt.Sprintf("Temperature %6.2f", obj.Temperature)
}