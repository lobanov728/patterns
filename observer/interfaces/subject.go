package interfaces

type Subject interface {
	RegisterObserver(item Observer)
	RemoveObserver(item Observer)
	NotifyObservers()
}
