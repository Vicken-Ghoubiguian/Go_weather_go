# Go_weather_go

## Contents

1. [Presentation](#presentation)

2. [How run the test programs](#how_run_the_test_programs)

<a name="presentation"></a>
## Presentation

This is a golang project to develop an API to get weather datas using the OpenWeatherMap utility.

Here is the datas:

* coordinates (longitude and latitude),

* main weather,

* weather description,

* current temperature,

* feeling temperature,

* maximum expected temperature,

* minimum expected temperature,

* sunrise time,

* sunset time,

* atmospheric pressure,

* atmospheric pressure at sea level,

* UV (index and risk),

* humidity,

* wind (speed and direction),

<a name="how_run_the_test_programs"></a>
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
