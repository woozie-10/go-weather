package tests

import (
	"encoding/json"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/woozie-10/go-weather/weather"
)

func TestGetCurrent(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	mockResponse := weather.CurrentWeather{
		Location: weather.Location{
			Name: "London",
		},
		Current: weather.Current{
			TempC: 25.0,
		},
	}
	mockResponseBody, _ := json.Marshal(mockResponse)

	httpmock.RegisterResponder(
		"GET",
		"http://api.weatherapi.com/v1/current.json?key=test_api_key&q=London",
		httpmock.NewStringResponder(200, string(mockResponseBody)),
	)

	wc := &weather.WeatherClient{API_KEY: "test_api_key"}

	weatherData, err := wc.GetCurrent("London")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if weatherData.Location.Name != "London" {
		t.Errorf("expected location name to be 'London', got '%s'", weatherData.Location.Name)
	}
	if weatherData.Current.TempC != 25.0 {
		t.Errorf("expected temperature to be 25.0, got %f", weatherData.Current.TempC)
	}
	t.Log("Test completed successfully")
}
