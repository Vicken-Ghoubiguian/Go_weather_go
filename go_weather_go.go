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

//The main function is the entry point of the go_weather_go utility
func main() {

	//presentationFunction is called to display the utility's presentation
	presentationFunction()

	//The parameters flags are defined below
	cityName := flag.String("city", "", "The city whose you want weather")
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")
	tempScale := flag.String("tempScale", "", "The temperature scale")

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

	var weather_map map[string]interface{}

	json.Unmarshal([]byte(weather_json_string), &weather_map)

	cod_as_string := fmt.Sprintf("%v", weather_map["cod"])

	code_as_int,_ := strconv.Atoi(cod_as_string)

	//If the returned code is different from 200 (the http request is successful)
	if code_as_int != 200 {

		//owmErrorHandler is called to display the occured http request error's code and message
		owmErrorHandler(cod_as_string, fmt.Sprintf("%v", weather_map["message"]))

	} else {

		fmt.Println("\n")

		fmt.Println("---------------------")

		fmt.Println("\n")

		fmt.Println("MAIN: ", weather_map["main"])
	}
}
