package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewWeatherClient(apiKey string) *WeatherClient {
	return &WeatherClient{
		API_KEY: apiKey,
	}
}
func (wc *WeatherClient) GetCurrent(city string) (CurrentWeather, error) {
	requestURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", wc.API_KEY, city)
	res, err := http.Get(requestURL)
	if err != nil {
		return CurrentWeather{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CurrentWeather{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	var weather CurrentWeather
	err = json.NewDecoder(res.Body).Decode(&weather)
	if err != nil {
		return CurrentWeather{}, err
	}

	return weather, nil

}


