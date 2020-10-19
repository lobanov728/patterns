package structs

import (
	"fmt"
	"patterns/observer/interfaces"
)

type SecondDisplay struct {
	Name string
	Humidity float64
}
func (obj *SecondDisplay) Update(subject interfaces.Subject) {
	fmt.Println("update second display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.Humidity = weatherData.Value2
	}
}
func (obj SecondDisplay) Display() string {
	return fmt.Sprintf("Humidity %6.2f", obj.Humidity)
}