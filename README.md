# Go_weather_go

## Contents

1. [Presentation](#presentation)

2. [How run the utility](#how_run_the_utility)

<a name="presentation"></a>
## Presentation

This is a project to develop a golang utility to get weather datas using the OpenWeatherMap API.

Here is the datas returned:

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

<a name="how_run_the_utility"></a>
## How run the utility ?

To run this program, you can execute this commands:

```bash
go build go_weather_go.go

./go_weather_go -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
or this single one:

```bash
go run go_weather_go.go -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
