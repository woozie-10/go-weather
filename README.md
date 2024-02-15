# Go Weather API Wrapper

This is a Go package that provides a wrapper for interacting with the [Weather API](https://www.weatherapi.com/).

## Installation
You can install the package using `go get`:
```bash
go get github.com/woozie-10/go-weather
```
## Usage
```go
package main

import (
	"fmt"

	"github.com/woozie-10/go-weather/weather"
)

func main() {
	client := weather.NewWeatherClient("your_api_key")
	// Get current weather
	current, err := client.GetCurrent("London")
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Println(current.Location.Country) // "United Kingdom"
	fmt.Println(current.Current.TempC)    // 12
	fmt.Println(current.Current.WindKph)  // 20.2
	/*---------------------------------------------------------------------*/
	// Get weather forecast
	forecast, err := client.GetForecast("London", 3)
	if err != nil {
		// Error handling
		panic(err)
	}
	fmt.Println(forecast.Forecast.Forecastday[0].Day.MaxtempC)       // 12.1
	fmt.Println(forecast.Forecast.Forecastday[1].Day.Condition.Text) // "Patchy rain nearby"
	fmt.Println(forecast.Forecast.Forecastday[2].Day.MaxwindKph)     // 19.8
}
```
## Testing
You can run tests using the `go test` command:
```bash
go test ./tests
```
## Contributing
If you have suggestions to improve this project, please open a new issue or pull request.

**Note:** Replace "YOUR_API_KEY" with your actual Weather API key.