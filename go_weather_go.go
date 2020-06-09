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

var red string = ""
var green string = ""
var reset string = ""

func owmErrorHandler(code_error string, error_message string) {

	fmt.Println("\n")

	fmt.Println(red + "Occured error (" + code_error + "): " + error_message + reset)
}

func errorHandlerFunction(err error) {
	
	if err != nil {

		fmt.Println(red + err + reset)

		os.Exit(1)
	}
}

func main() {

	cityName := flag.String("city", "", "The city whose you want weather")

	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")

	flag.Parse()

	weather_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *cityName, *apiKey)

	resp, err := http.Get(weather_request)

	errorHandlerFunction(err)

	weather_json_string, err := ioutil.ReadAll(resp.Body)

	errorHandlerFunction(err)

	fmt.Println(string(weather_json_string))

	var weather_map map[string]interface{}

	json.Unmarshal([]byte(weather_json_string), &weather_map)
	
	cod_as_string := fmt.Sprintf("%v", weather_map["cod"])

	error_nbr,_ := strconv.Atoi(cod_as_string)

	if error_nbr != 200 {
		
		owmErrorHandler(cod_as_string, fmt.Sprintf("%v", weather_map["message"]))

	} else {

		fmt.Println("\n")

		fmt.Println("---------------------")

		fmt.Println("\n")

		fmt.Println("MAIN: ", weather_map["main"])
	}
}
