package observer

import "testing"
import "strings"

func TestWeatherStations(t *testing.T) {
	temperature := "10"
	pressure := "20"

	weatherData := WeatherData{}
	weatherData.SetTemerature(temperature)
	weatherData.SetPressure(pressure)

	tempObserver := TemperatureObserver{}
	pressObserver := PressureObserver{}

	weatherStation := WeatherStation{}
	weatherStation.AddObserver(&tempObserver)
	weatherStation.AddObserver(&pressObserver)
	weatherStation.SetWeatherData(&weatherData)

	var checkData = func() {
		if currentTemperature := tempObserver.display(); !strings.Contains(currentTemperature, temperature) {
			t.Errorf("Wrong temperature. Expected %s, got %s", temperature, currentTemperature)
		}

		if currentPressure := pressObserver.display(); !strings.Contains(currentPressure, pressure) {
			t.Errorf("Wrong pressure. Expected %s, got %s", pressure, currentPressure)
		}
	}

	checkData()

	weatherStation.RemoveObserver(&tempObserver)
	weatherStation.RemoveObserver(&pressObserver)

	weatherData.SetTemerature("90")
	weatherData.SetPressure("50")

	weatherStation.SetWeatherData(&weatherData)

	checkData()
}
