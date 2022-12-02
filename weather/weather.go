package weather

import (
    "encoding/json"
    "fmt"
    "fraser-chapman/go-api/error"
    "io"
    "net/http"
)
type CurrentWeather struct {
	Time        string  `json:"time"`
	Temperature float32 `json:"temperature"`
}
type Weather struct {
	CurrentWeather CurrentWeather `json:"current_weather"`
}

func Get(lat float64, long float64) Weather {
    url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", lat, long)

    client := http.Client{}
    response, err := client.Get(url)
    error.Handle(err)

    bodyBytes, err := io.ReadAll(response.Body)
    error.Handle(err)

    var weather Weather
    err = json.Unmarshal(bodyBytes, &weather)
    error.Handle(err)

    return weather
}
