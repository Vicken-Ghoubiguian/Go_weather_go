package main
import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
    "flag"
)

func main() {

	cityName := flag.String("city", "", "The city whose you want weather")

	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")

	flag.Parse()

	weather_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *cityName, *apiKey)

	resp, err := http.Get(weather_request)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	text, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Println(string(text))
}
