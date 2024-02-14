package main

import (
	"fmt"

	"github.com/woozie-10/go-weather/weather"
)

func main() {
	client := weather.NewWeatherClient("07e78758e45c43ff8d790211241402")
	// Get current weather
	current, err := client.GetCurrent("London")
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Println(current.Location.Country) // "United Kingdom"
	fmt.Println(current.Current.TempC) // 12
	fmt.Println(current.Current.WindKph) // 20.2
	/*---------------------------------------------------------------------*/
	// Get weather forecast
	forecast, err := client.GetForecast("London", 3)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Println(forecast.Forecast.Forecastday[0].Day.MaxtempC) // 12.1
	fmt.Println(forecast.Forecast.Forecastday[1].Day.Condition.Text) // "Patchy rain nearby"
	fmt.Println(forecast.Forecast.Forecastday[2].Day.MaxwindKph) // 19.8
}
