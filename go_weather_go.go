package main

//Import all necessary packages
import (
    "net/http"
    "strconv"
    "fmt"
    "encoding/json"
    "io/ioutil"
    "os"
    "flag"
)

//Definition of global variables which represents colors
var red string = "\033[31m"
var green string = "\033[32m"
var cyan string = "\033[36m"
var reset string = "\033[0m"

//Function which display the utility's presentation
func presentationFunction() {

	fmt.Println(cyan + "Welcome to Go weather go !!!!" + reset)

	fmt.Println(cyan + "This is a simple client to display current weather at your favorite town from OpenWeatherMap API" + reset)

	fmt.Println("\n")
}

//Function which display http request error's code and message when the first occurs
func owmErrorHandler(code_error string, error_message string) {

	fmt.Println(red + "Occured error (" + code_error + "): " + error_message + reset)

	fmt.Println("\n")
}

//Function which display other general errors when they occurs
func errorHandlerFunction(err error) {

	if err != nil {

		fmt.Println(red + err.Error() + reset)

		os.Exit(1)
	}
}

//Function which converts the kelvin temperature in the specified temperature scale
func temperatureConversionFunction(temperatureInKelvin float64, temperatureScale *string) float64 {

	var convertedTemperatureValue float64 = 0.0
	var tempScaleAsString string = *temperatureScale

	if tempScaleAsString == "celsius" {

		convertedTemperatureValue = temperatureInKelvin - 273.15

	} else if tempScaleAsString == "fahrenheit" {

		convertedTemperatureValue = temperatureInKelvin * (9/5) - 459.67

	} else if tempScaleAsString == "kelvin" {

		convertedTemperatureValue = temperatureInKelvin

	} else {

		fmt.Println(red + "Error: " + tempScaleAsString + " is not a temperature scale" + reset)

		os.Exit(1)

	}

	return convertedTemperatureValue
}

type cod struct {

	COD int `json:"cod"`
}

type message struct {

	MESSAGE string `json:"message"`
}

//The main function is the entry point of the go_weather_go utility
func main() {

	//presentationFunction is called to display the utility's presentation
	presentationFunction()

	//The parameters flags are defined below
	cityName := flag.String("city", "", "The city whose you want weather")
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")
	tempScale := flag.String("tempScale", "kelvin", "The temperature scale")

	_ = temperatureConversionFunction(0.0, tempScale)

	//Parsing all received values for each flag
	flag.Parse()

	//Definition of the URL for the http request string and affectation of it in the weather_request variable
	weather_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *cityName, *apiKey)

	//Make the http get request and affectation of it's response in the the resp variable as JSON string
	resp, err := http.Get(weather_request)

	//errorHandlerFunction is called to treat any occured error from the above instruction
	errorHandlerFunction(err)

	//
	weather_json_string, err := ioutil.ReadAll(resp.Body)

	//errorHandlerFunction is called to treat any occured error from the above instruction
	errorHandlerFunction(err)

	//Single instruction for testing and development
	//fmt.Println(string(weather_json_string))

	//Declaration of struct instances
	message1 := message{}
	cod1 := cod{}

	var weather_map map[string]interface{}

	json.Unmarshal([]byte(weather_json_string), &weather_map)

	json.Unmarshal([]byte(weather_json_string), &cod1)

	json.Unmarshal([]byte(weather_json_string), &message1)

	//If the returned code is different from 200 (the http request is successful)
	if cod1.COD != 200 {

		//owmErrorHandler is called to display the occured http request error's code and message
		owmErrorHandler(strconv.Itoa(cod1.COD), message1.MESSAGE)

	} else {

		fmt.Println("\n")

		fmt.Println("---------------------")

		fmt.Println("\n")

		fmt.Println("MAIN: ", weather_map["main"])
	}
}
