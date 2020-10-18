package main

import "fmt"

type Observer interface {
	Update(subject Subject) // в update закидываем объект
}

type Subject interface {
	registerObserver(item Observer)
	removeObserver(item Observer)
	notifyObservers()
}

type WeatherData struct { // реализует интерфейс Subject
	observers []Observer // слайс указателей на объекты Observer
	value1, value2, value3 float64
}
func (weatherData *WeatherData) registerObserver(item Observer) {
	fmt.Println(fmt.Sprintf("type of item %T", item))
	weatherData.observers = append(weatherData.observers, item)
	fmt.Println("observers slice content", weatherData.observers)
}

func (weatherData *WeatherData) removeObserver(item Observer) {
	key := -1
	for i, existObserver := range weatherData.observers {
		if existObserver == item {
			key = i
			break
		}
	}
	if key > -1 {
		weatherData.observers = append(weatherData.observers[:key], weatherData.observers[key+1:]...)
	}

}
func (weatherData WeatherData) notifyObservers() {
	fmt.Println("Notify observers")
	for _, item := range weatherData.observers {
		fmt.Println(fmt.Sprintf("Observer type %T", item))
		(item).Update(&weatherData)
	}
}

type FirstDisplay struct {
	name              string
	Temperature       float64
}
type SecondDisplay struct {
	name string
	humidity float64
}
type ThirdDisplay struct {
	name string
	velocity float64
}

func (obj *FirstDisplay) Update(subject Subject) {
	fmt.Println("update first display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.Temperature = weatherData.value1
	}
}
func (obj FirstDisplay) display() {
	fmt.Println("Temperature", obj.Temperature)
}
func (obj *SecondDisplay) Update(subject Subject) {
	fmt.Println("update second display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.humidity = weatherData.value2
	}
}
func (obj SecondDisplay) display() {
	fmt.Println("Humidity", obj.humidity)
}
func (obj *ThirdDisplay) Update(subject Subject) {
	fmt.Println("update third display called")
	fmt.Printf("%T", subject)
	weatherData, ok := subject.(*WeatherData)
	fmt.Println("OK", ok)
	if ok {
		obj.velocity = weatherData.value3 * 3
	}
}
func (obj ThirdDisplay) display() {
	fmt.Println("Velocity", obj.velocity)
}

func RunObserver() {
	weatherData := WeatherData{
		observers: nil,
		value1:    0,
		value2:    0,
		value3:    0,
	}

	firstDisplay := FirstDisplay{name: "temperature"}
	secondDisplay := &SecondDisplay{name: "humidity"}
	thirdDisplay := &ThirdDisplay{name: "velocity"}
	observers := []Observer{
		&firstDisplay,
		secondDisplay,
		thirdDisplay,
	}
	for i, observer := range observers {
		fmt.Println(i)
		weatherData.registerObserver(observer)
	}

	weatherData.value1 = 10008
	weatherData.value2 = 20008
	weatherData.value3 = 30008
	weatherData.notifyObservers()
	firstDisplay.display()
	secondDisplay.display()
	thirdDisplay.display()

	weatherData.removeObserver(observers[0])
	weatherData.value1 = 100081
	weatherData.value2 = 200081
	weatherData.value3 = 300081
	weatherData.notifyObservers()
	firstDisplay.display()
	secondDisplay.display()
	thirdDisplay.display()
}

func main() {
	fmt.Println("Observer")
	RunObserver()
}