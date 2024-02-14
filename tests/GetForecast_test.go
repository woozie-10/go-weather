package tests

import (
	"testing"

	"github.com/woozie-10/go-weather/weather"
)

func TestGetForecast(t *testing.T) {
	client := weather.NewWeatherClient("test_api_key")
	res, err := client.GetForecast("London", 5)
	if err != nil {
		t.Errorf("ERROR: %v", err)
	}
	if res.Location.Name != "London" {
		t.Error("ERROR: the result does not match")
	}
	t.Log("Test completed successfully")

}
