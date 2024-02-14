package main

import (
	"fmt"

	"github.com/woozie-10/go-weather/weather"
)

func main() {
	client := weather.NewWeatherClient("07e78758e45c43ff8d790211241402")
	currentWeather, err := client.GetCurrent("Kharkiv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(currentWeather.Location.Name)
}
