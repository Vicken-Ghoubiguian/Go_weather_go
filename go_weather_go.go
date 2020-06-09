package main

import (
    "net/http"
    "strconv"
    "fmt"
    "encoding/json"
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

	weather_json_string, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		fmt.Println(err)

		os.Exit(1)

	}

	fmt.Println(string(weather_json_string))

	var weather_map map[string]interface{}

	json.Unmarshal([]byte(weather_json_string), &weather_map)
	
	cod_as_string := fmt.Sprintf("%v", weather_map["cod"])

	error_nbr,_ := strconv.Atoi(cod_as_string)

	if error_nbr != 200 {
	
		error_message := fmt.Sprintf("%v", weather_map["message"])

		fmt.Println("\n")

		fmt.Println("Occured error (" + cod_as_string + "): " + error_message)

	} else {

		fmt.Println("\n")

		fmt.Println("---------------------")

		fmt.Println("\n")

		fmt.Println("MAIN: ", weather_map["main"])
	}
}
