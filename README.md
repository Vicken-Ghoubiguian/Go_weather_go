# Go_weather_go

This is a golang project to develop an API to get weather datas using the OpenWeatherMap utility.



## How run the test programs ?

All test programs are in the `tests` folder.

So first of all, you must position yourself in the `tests` folder of the project by executing this command where you cloned this repo:

```bash
cd go_weather_go/tests
```
Now to run the test program you want, you can execute this commands:

```bash
go build <wished_test_program_name>.go

./<wished_test_program_name> -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
or this single one:

```bash
go run <wished_test_program_name>.go -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
