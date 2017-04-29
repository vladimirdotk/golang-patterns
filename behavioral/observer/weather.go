package observer

import (
	"fmt"
)

const tempFormat = "The temperature is %s\n"
const presFormat = "The pressure is %s\n"

type WeatherStation struct {
	data      WeatherData
	observers []*WeatherObserver
}

func (ws *WeatherStation) SetWeatherData(data *WeatherData) {
	ws.data = *data
	ws.notifyObservers()
}

func (ws *WeatherStation) notifyObservers() {
	for _, observer := range ws.observers {
		(*observer).update(ws.data)
	}
}

func (ws *WeatherStation) AddObserver(observer WeatherObserver) bool {
	ok := false
	if _, exsists := ws.isObserverExists(observer); !exsists {
		ws.observers = append(ws.observers, &observer)
		ok = true
	}
	return ok
}

func (ws *WeatherStation) RemoveObserver(observer WeatherObserver) bool {
	ok := false
	if index, exsists := ws.isObserverExists(observer); exsists {
		ws.deleteObserver(index)
		ok = true
	}
	return ok
}

func (ws *WeatherStation) deleteObserver(index int) {
	copy(ws.observers[index:], ws.observers[index+1:])
	ws.observers[len(ws.observers)-1] = nil
	ws.observers = ws.observers[:len(ws.observers)-1]
}

func (ws *WeatherStation) isObserverExists(observer WeatherObserver) (int, bool) {
	for index, element := range ws.observers {
		if *element == observer {
			return index, true
		}
	}
	return -1, false
}

type WeatherData struct {
	temperature string
	pressure    string
}

func (wd *WeatherData) SetTemerature(temperature string) {
	wd.temperature = temperature
}

func (wd *WeatherData) GetTemperature() string {
	return wd.temperature
}

func (wd *WeatherData) SetPressure(pressure string) {
	wd.pressure = pressure
}

func (wd *WeatherData) GetPressure() string {
	return wd.pressure
}

type WeatherObserver interface {
	update(WeatherData)
	display() string
}

type TemperatureObserver struct {
	temperature string
}

func (to *TemperatureObserver) update(data WeatherData) {
	to.temperature = data.GetTemperature()
}

func (to *TemperatureObserver) display() string {
	return fmt.Sprintf(tempFormat, to.temperature)
}

type PressureObserver struct {
	pressure string
}

func (po *PressureObserver) update(data WeatherData) {
	po.pressure = data.pressure
}

func (po *PressureObserver) display() string {
	return fmt.Sprintf(presFormat, po.pressure)
}
