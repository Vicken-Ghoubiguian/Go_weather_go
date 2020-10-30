package main

//Import all necessary packages
import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
    "flag"
    "time"
    "strconv"
    "encoding/json"
)

type Coordinates struct {
	Lon float64
	Lat float64
}

type Clouds struct {
	All int
}

type Wind struct {
	Speed float64
	Deg int
}

type Main struct {
	Temp float64
	Feels_like float64
	Temp_min float64
	Temp_max float64
	Pressure int
	Humidity int
}

type Weather struct {
	Id int
	Main string
	Description string
	Icon string
}

type Sys struct {
	Type int
	Id int
	Country string
	Sunrise int
	Sunset int
}

type UVStruct struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
	Date_iso string `json:"date_iso"`
	Date int `json:"date"`
	Value float64 `json:"value"`
}

type OWMStruct struct {
	Coord Coordinates `json:"coord"`
	Weather [1]Weather `json:"weather"`
	Base string `json:"base"`
	Main Main `json:"main"`
	Visibility int `json:"visibility"`
	Wind Wind `json:"wind"`
	Clouds Clouds `json:"clouds"`
	Dt int `json:"dt"`
	Sys Sys `json:"sys"`
	Timezone int `json:"timezone"`
	Id int `json:"id"`
	Name string `json:"name"`
	Cod int `json:"cod"`
	Message string `json:"message"`
}

//Importation of the github project gjson to treat received json
//import "github.com/tidwall/gjson"

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

		convertedTemperatureValue = temperatureInKelvin * 1.8 - 459.67

	} else if tempScaleAsString == "kelvin" {

		convertedTemperatureValue = temperatureInKelvin

	} else {

		fmt.Println(red + "Error: " + tempScaleAsString + " is not a temperature scale" + reset)

		os.Exit(1)

	}

	return convertedTemperatureValue
}

//Function which defines the temperature scale symbol from the specified temperature scale
func temperatureSymbolFunction(temperatureScale *string) string {

	var returnedSymbol string = ""
	var tempScaleAsString string = *temperatureScale

	if tempScaleAsString == "celsius" {

                returnedSymbol = " °C"

        } else if tempScaleAsString == "fahrenheit" {

                returnedSymbol = " °F"

        } else if tempScaleAsString == "kelvin" {

                returnedSymbol = " K"

        } else {

                fmt.Println(red + "Error: " + tempScaleAsString + " is not a temperature scale" + reset)

                os.Exit(1)

        }

        return returnedSymbol
}

//Function which extracts weather datas from JSON response
/*func extractWeatherFromJSONFunction(weather_from_http_response string) string {

	brut_weather_without_hooks := strings.Trim(weather_from_http_response, "[]")

	return brut_weather_without_hooks
}*/

//Function to convert timestamp to readable and formated time
func treatingAndFormatingFunction(time_as_timestamp int) string {

	time_as_time := time.Unix(int64(time_as_timestamp), 0)

	time_as_string := time_as_time.Format(time.UnixDate)

	return time_as_string
}

//Function to determine current UV risk level
func riskDeterminationFunction(uvValue int) string {

	var uvRiskValue string = ""

	if uvValue <= 2 && 0 <= uvValue {

		uvRiskValue = "Low"

	} else if uvValue <= 5 && 3 <= uvValue {

		uvRiskValue = "Moderate"

	} else if uvValue <= 7 && 6 <= uvValue {

		uvRiskValue = "High"

	} else if uvValue <= 10 && 8 <= uvValue {

		uvRiskValue = "Very high"

	} else if 11 <= uvValue {

		uvRiskValue = "Extreme"

	} else {

		fmt.Println(red + "Error: " + "not available value for UV index" + reset)

		os.Exit(1)

	}

	return uvRiskValue
}

//The main function is the entry point of the go_weather_go utility
func main() {

	//
	var owm OWMStruct

	//presentationFunction is called to display the utility's presentation
	presentationFunction()

	//The parameters flags are defined below
	cityName := flag.String("city", "", "The city whose you want weather")
	apiKey := flag.String("apiKey", "", "The OpenWeatherMap API key")
	tempScale := flag.String("tempScale", "kelvin", "The temperature scale")

	//Parsing all received values for each flag
	flag.Parse()

	//Definition of the URL for the http request string and affectation of it in the weather_request variable
	weather_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", *cityName, *apiKey)

	//Make the http get request and affectation of it's response in the the resp variable as JSON string
	resp, err := http.Get(weather_request)

	//errorHandlerFunction is called to treat any occured error from the above instruction
	errorHandlerFunction(err)

	//
	weather_json, err := ioutil.ReadAll(resp.Body)

	//errorHandlerFunction is called to treat any occured error from the above instruction
	errorHandlerFunction(err)

	//Single instruction to convert weather_json_string []byte variable to string
	//weather_string := string(weather_json_string)
	err = json.Unmarshal(weather_json, &owm)

	//errorHandlerFunction is called to treat any occured error from the above instruction
	errorHandlerFunction(err)

	//Single instruction for testing and development
	//fmt.Println(weather_string)

	//Extraction of HTTP code
	//code := gjson.Get(weather_string, "cod")

	//If the returned code is different from 200 (the http request is successful)
	if owm.Cod != 200 {

		//owmErrorHandler is called to display the occured http request error's code and message
		owmErrorHandler(strconv.Itoa(owm.Cod), owm.Message)

	} else {

		//
		var UVowm UVStruct

		//Breaking the line
		fmt.Println("\n")

		//Extraction of name and country code from the JSON string in the variable 'weather_string'
		/*cityName := gjson.Get(weather_string, "name")
		countryCode := gjson.Get(weather_string, "sys.country")*/

		/*cityName := owm.Name
		countryCode := owm.Sys.Country*/

		//Extraction of longitude and latitude from the JSON string in the variable 'weather_string'
		/*longitude := owm.Coordinates.Lon
		latitude := owm.Coordinates.Lat*/

		//Definition of the URL for the http request string and affectation of it in the uvi_request variable
		uvi_request := fmt.Sprintf("https://api.openweathermap.org/data/2.5/uvi?appid=%s&lat=%s&lon=%s", *apiKey, fmt.Sprintf("%g", owm.Coord.Lat), fmt.Sprintf("%g", owm.Coord.Lon))

		//Make the http get request and affectation of it's response in the the resp variable as JSON string
		resp, err := http.Get(uvi_request)

		//errorHandlerFunction is called to treat any occured error from the above instruction
		errorHandlerFunction(err)

		//
		uvi_json, err := ioutil.ReadAll(resp.Body)

		//errorHandlerFunction is called to treat any occured error from the above instruction
		errorHandlerFunction(err)

		//Single instruction to convert uvi_json_string []byte variable to string
		//uvi_string := string(uvi_json)

		//Single instruction for testing and development
		//fmt.Println(uvi_string)

		err = json.Unmarshal(uvi_json, &UVowm)

		//errorHandlerFunction is called to treat any occured error from the above instruction
		errorHandlerFunction(err)

		//Extraction of the uv index value from the JSON string variable uvi_string
		/*uvi_brut_value := gjson.Get(uvi_string, "value")

		//Extraction of the uv index's measure date from the JSON string variable uvi_string
		uvi_measure_date := gjson.Get(uvi_string, "date")

		//Conversion of the uv index's measure date to int
		uvi_measure_date_as_timestamp := uvi_measure_date.Int()

		//Rounding UV index value and store it to the uvi_value variable
		uvi_value := uvi_brut_value.Int()*/

		//Extraction of weather datas from JSON obtained from HTTP response
		//weather := extractWeatherFromJSONFunction(gjson.Get(weather_string, "weather").String())

		//
		//uvi_measure_date_as_timestamp := UVowm.Date

		//
		//uvi_value := UVowm.Value

		//Extraction of all weather datas from JSON string in the variable 'weather'
		/*main_weather := gjson.Get(weather, "main")
		description_weather := gjson.Get(weather, "description")*/

		main_weather := owm.Weather[0].Main
		description_weather := owm.Weather[0].Description

		//Extraction of all datas relative to temperature from the JSON string in the variable 'weather_string'
		/*temperature := gjson.Get(weather_string, "main.temp")
		feeling_temperature := gjson.Get(weather_string, "main.feels_like")
		minimum_temperature := gjson.Get(weather_string, "main.temp_min")
		maximum_temperature := gjson.Get(weather_string, "main.temp_max")*/

		//Extraction of all datas relative to wind from the JSON variable weather_string
		/*wind_speed := gjson.Get(weather_string, "wind.speed")
		wind_direction := gjson.Get(weather_string, "wind.deg")

		//Extraction of humidity from the JSON variable weather_string to the humidity variable
		humidity := gjson.Get(weather_string, "main.humidity")

		//Extraction of atmospheric pressure from the JSON variable weather_string to the atmospheric_pressure variable
		atmospheric_pressure := gjson.Get(weather_string, "main.pressure")*/

		//Displaying wished city and the corresponding country code
		fmt.Println(green + owm.Name + " (" + owm.Sys.Country + ")" + reset)

		//Displaying of all weather elements
		fmt.Println(green + "Geographic coordinates: (longitude: ", fmt.Sprintf("%g", owm.Coord.Lon), ", latitude: ", fmt.Sprintf("%g", owm.Coord.Lat), ")" + reset)
		fmt.Println(green + "Main weather: ", main_weather, "" + reset)
		fmt.Println(green + "Description weather: ", description_weather, "" + reset)

		//Convert all datas about temperature in the wished scale
		/*temperature_in_wished_scale := temperatureConversionFunction(temperature.Float(), tempScale)
		feeling_temperature_in_wished_scale := temperatureConversionFunction(feeling_temperature.Float(), tempScale)
		minimum_temperature_in_wished_scale := temperatureConversionFunction(minimum_temperature.Float(), tempScale)
		maximum_temperature_in_wished_scale := temperatureConversionFunction(maximum_temperature.Float(), tempScale)*/

		temperature_in_wished_scale := temperatureConversionFunction(owm.Main.Temp, tempScale)
		feeling_temperature_in_wished_scale := temperatureConversionFunction(owm.Main.Feels_like, tempScale)
		minimum_temperature_in_wished_scale := temperatureConversionFunction(owm.Main.Temp_min, tempScale)
		maximum_temperature_in_wished_scale := temperatureConversionFunction(owm.Main.Temp_max, tempScale)

		//Define the temperature scale symbol
		temperatureScaleSymbol := temperatureSymbolFunction(tempScale)

		//Displaying of all datas about temperature
		fmt.Println(green + "Temperature: ", fmt.Sprintf("%.2f", temperature_in_wished_scale), temperatureScaleSymbol + reset)
		fmt.Println(green + "Feeling temperature: ", fmt.Sprintf("%.2f", feeling_temperature_in_wished_scale), temperatureScaleSymbol + reset)
		fmt.Println(green + "Minimum temperature: ", fmt.Sprintf("%.2f", minimum_temperature_in_wished_scale), temperatureScaleSymbol + reset)
		fmt.Println(green + "Maximum temperature: ", fmt.Sprintf("%.2f", maximum_temperature_in_wished_scale), temperatureScaleSymbol + reset)

		//Extraction of sunrise and sunset time as timestamp
		/*sunrise := gjson.Get(weather_string, "sys.sunrise")
		sunset := gjson.Get(weather_string, "sys.sunset")*/

		sunrise := owm.Sys.Sunrise
		sunset := owm.Sys.Sunset

		//Conversion of collected datas to int
		/*sunrise_as_timestamp := sunrise.Int()
		sunset_as_timestamp := sunset.Int()*/

		//Displaying sunrise and sunset time
		fmt.Println(green + "Sunrise: ", treatingAndFormatingFunction(sunrise), reset)
		fmt.Println(green + "Sunset: ", treatingAndFormatingFunction(sunset), reset)

		//Displaying wind's speed and direction
		/*fmt.Println(green + "Wind speed: ", wind_speed.String(), " m/s" + reset)
		fmt.Println(green + "Wind direction: ", wind_direction.String(), " °" + reset)*/
		fmt.Println(green + "Wind speed: ", fmt.Sprintf("%g", owm.Wind.Speed), " m/s" + reset)
		fmt.Println(green + "Wind direction: ", strconv.Itoa(owm.Wind.Deg), " °" + reset)

		//Displaying humidity
		//fmt.Println(green + "Humidity: ", humidity.String(), " %" + reset)
		fmt.Println(green + "Humidity: ", strconv.Itoa(owm.Main.Humidity), " %" + reset)

		//Displaying atmospheric pressure
		//fmt.Println(green + "Atmospheric pressure: ", atmospheric_pressure.String(), " hPa" + reset)
		fmt.Println(green + "Atmospheric pressure: ", strconv.Itoa(owm.Main.Pressure), " hPa" + reset)

		//Displaying all necessary datas about UV index
		fmt.Println(green + "UV index (at ", treatingAndFormatingFunction(UVowm.Date), "): ", strconv.Itoa(int(int64(UVowm.Value))), "(risk: " + riskDeterminationFunction(int(int64(UVowm.Value))) + ")" + reset)

		//Breaking another line
		fmt.Println("\n")
	}
}
