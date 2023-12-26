package weather

import (
	"slices"

	log "github.com/sirupsen/logrus"
)

// types
type WeatherData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
	observers   []Observer
}

func (w *WeatherData) MeasurementsChagned() {
	log.Infoln("measurements changed")
	w.NotifyObservers()
}

type CurrentConditionsDisplay struct {
	Wd          *WeatherData
	Temperature float64
	Humidity    float64
}

func (c *CurrentConditionsDisplay) Update() {
	c.Temperature = c.Wd.Temperature
	c.Humidity = c.Wd.Humidity
	c.Display()
}

func (c *CurrentConditionsDisplay) Display() {
	log.Infof("Temperature: %f, Humidity: %f\n", c.Temperature, c.Humidity)
}

// interfaces
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Observer interface {
	Update()
}

type DisplayElement interface {
	Display()
}

func (w *WeatherData) RegisterObserver(o Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	if idx := slices.Index(w.observers, o); idx >= 0 {
		w.observers = append(w.observers[:idx], w.observers[idx+1:]...)
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, o := range w.observers {
		o.Update()
	}
}

// methods
func (w *WeatherData) setMeasurements(
	temperature float64, humidity float64, pressure float64) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure
	w.MeasurementsChagned()
}

// functions
func Run() {
	log.Infoln("--- Weather ---")
	wd := &WeatherData{}

	c := &CurrentConditionsDisplay{
		Wd: wd,
	}
	wd.RegisterObserver(c)
	c2 := &CurrentConditionsDisplay{
		Wd: wd,
	}
	wd.RegisterObserver(c2)
	c3 := &CurrentConditionsDisplay{
		Wd: wd,
	}
	wd.RegisterObserver(c3)

	wd.setMeasurements(1.0, 2.0, 3.0)
	wd.RemoveObserver(c3)

	wd.setMeasurements(11.0, 12.0, 13.0)
	wd.RemoveObserver(c2)

	wd.setMeasurements(21.0, 22.0, 23.0)
}
