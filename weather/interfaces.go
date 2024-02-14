package weather

type WeatherClientInterface interface {
	GetCurrent(city string) (CurrentWeather, error)
	GetForecast(city string, days int) (ForecastWeather, error)
}



