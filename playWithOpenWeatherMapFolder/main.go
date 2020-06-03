package main
import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

	weather_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", "<city_name>", "<api_key>")

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
