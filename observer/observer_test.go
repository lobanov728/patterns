package main

import (
	"patterns/observer/interfaces"
	"patterns/observer/structs"
	"reflect"
	"strconv"
	"testing"
)

func TestRegisterObserverNotEmpty(t *testing.T) {
	weatherData := structs.WeatherData{
		Observers: nil,
		Value1:    0,
		Value2:    0,
		Value3:    0,
	}
	firstDisplay := structs.FirstDisplay{Name: "test"}
	observers := []interfaces.Observer {
		&firstDisplay,
	}
	for _, observer := range observers {
		weatherData.RegisterObserver(observer)
	}
	if !reflect.DeepEqual(1, len(weatherData.Observers)) {
		t.Error("Weather data has not proper count of observers")
	}
}

func TestRemoveObserver(t *testing.T) {
	firstDisplay := structs.FirstDisplay{Name: "test1"}
	secondDisplay := structs.SecondDisplay{Name: "test2"}
	thirdDisplay := structs.ThirdDisplay{Name: "test3"}
	weatherData := structs.WeatherData{
		Observers: []interfaces.Observer {
			&firstDisplay,
			&secondDisplay,
			&thirdDisplay,
		},
		Value1:    0,
		Value2:    0,
		Value3:    0,
	}

	weatherData.RemoveObserver(&secondDisplay)
	if !reflect.DeepEqual(2, len(weatherData.Observers)) {
		t.Error("Weather data has not proper count of observers")
	}
}

func TestMainLogic(t *testing.T) {
	firstDisplay := structs.FirstDisplay{Name: "test1"}
	secondDisplay := structs.SecondDisplay{Name: "test2"}
	thirdDisplay := structs.ThirdDisplay{Name: "test3"}
	weatherData := structs.WeatherData{
		Observers: nil,
		Value1:    0,
		Value2:    0,
		Value3:    0,
	}
	observers := []interfaces.Observer {
		&firstDisplay,
		&secondDisplay,
		&thirdDisplay,
	}
	for _, observer := range observers {
		weatherData.RegisterObserver(observer)
	}
	if !reflect.DeepEqual(3, len(weatherData.Observers)) {
		t.Error("Weather data has not proper count of observers")
	}
	weatherData.Value1 = 10008
	weatherData.Value2 = 20008
	weatherData.Value3 = 30008
	weatherData.NotifyObservers()
	if !reflect.DeepEqual("Temperature 10008.00", firstDisplay.Display()) {
		t.Error("firstDisplay wrong display")
	}
	if !reflect.DeepEqual("Humidity 20008.00", secondDisplay.Display()) {
		t.Error("secondDisplay wrong display")
	}
	if !reflect.DeepEqual("Velocity " + strconv.Itoa(30008 * 3) + ".00", thirdDisplay.Display()) {
		t.Error("thirdDisplay wrong display")
	}

	weatherData.RemoveObserver(observers[0])
	weatherData.Value1 = 100081
	weatherData.Value2 = 200081
	weatherData.Value3 = 300081
	weatherData.NotifyObservers()
	firstDisplay.Display()
	secondDisplay.Display()
	thirdDisplay.Display()

	if !reflect.DeepEqual("Temperature 10008.00", firstDisplay.Display()) {
		t.Error("firstDisplay wrong display")
	}
	if !reflect.DeepEqual("Humidity 200081.00", secondDisplay.Display()) {
		t.Error("secondDisplay wrong display")
	}
	if !reflect.DeepEqual("Velocity " + strconv.Itoa(300081 * 3) + ".00", thirdDisplay.Display()) {
		t.Error("thirdDisplay wrong display")
	}
}