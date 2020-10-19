package structs

import "fmt"
import "patterns/observer/interfaces"

type WeatherData struct {
	Observers []interfaces.Observer
	Value1, Value2, Value3 float64
}
func (weatherData *WeatherData) RegisterObserver(item interfaces.Observer) {
	fmt.Println(fmt.Sprintf("type of item %T", item))
	weatherData.Observers = append(weatherData.Observers, item)
	fmt.Println("observers slice content", weatherData.Observers)
}

func (weatherData *WeatherData) RemoveObserver(item interfaces.Observer) {
	key := -1
	for i, existObserver := range weatherData.Observers {
		if existObserver == item {
			key = i
			break
		}
	}
	if key > -1 {
		weatherData.Observers = append(weatherData.Observers[:key], weatherData.Observers[key+1:]...)
	}
}
func (weatherData WeatherData) NotifyObservers() {
	fmt.Println("Notify observers")
	for _, item := range weatherData.Observers {
		fmt.Println(fmt.Sprintf("Observer type %T", item))
		(item).Update(&weatherData)
	}
}
