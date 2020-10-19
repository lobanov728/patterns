package main

import (
	"fmt"
	"patterns/observer/interfaces"
	"patterns/observer/structs"
)

func RunObserver() {
	weatherData := structs.WeatherData{
		Observers: nil,
		Value1:    0,
		Value2:    0,
		Value3:    0,
	}

	firstDisplay := structs.FirstDisplay{Name: "temperature"}
	secondDisplay := &structs.SecondDisplay{Name: "humidity"}
	thirdDisplay := &structs.ThirdDisplay{Name: "velocity"}
	observers := []interfaces.Observer{
		&firstDisplay,
		secondDisplay,
		thirdDisplay,
	}
	for i, observer := range observers {
		fmt.Println(i)
		weatherData.RegisterObserver(observer)
	}

	weatherData.Value1 = 10008
	weatherData.Value2 = 20008
	weatherData.Value3 = 30008
	weatherData.NotifyObservers()
	firstDisplay.Display()
	secondDisplay.Display()
	thirdDisplay.Display()

	weatherData.RemoveObserver(observers[0])
	weatherData.Value1 = 100081
	weatherData.Value2 = 200081
	weatherData.Value3 = 300081
	weatherData.NotifyObservers()
	firstDisplay.Display()
	secondDisplay.Display()
	thirdDisplay.Display()
}

func main() {
	fmt.Println("Observer")
	RunObserver()
}