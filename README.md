# Go_weather_go

## Contents

1. [Presentation](#presentation)

2. [Prerequisites to run the utility on machine host](#prerequisites_to_run_the_utility_on_machine_host)

3. [Prerequisites to run the utility with Docker](#prerequisites_to_run_the_utility_with_docker)

4. [How run the utility on machine host](#how_run_the_utility_on_machine_host)

5. [How run the utility with Docker](#how_run_the_utility_with_docker)

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

* UV index,

* humidity,

* wind (speed and direction),

<a name="prerequisites_to_run_the_utility_on_machine_host"></a>
## Prerequisites to run the utility on machine host

First, you must install Go language on your machine.

To do this, you can consult the official documentation [here](https://golang.org/doc/install).

Or you can install it by executing this command if you're on Debian or Ubuntu distribution:

```bash
sudo apt install golang-go
```
Finally, you must install the `gjson` github project to manipulate the JSON response from HTTP request.

To do this, execute the following command:

```bash
go get -u github.com/tidwall/gjson
```
Well done. It's your turn to play now...

<a name="prerequisites_to_run_the_utility_with_docker"></a>
## Prerequisites to run the utility with Docker

First, you must install Docker on your machine.

To do this, you can consult the official documentation [here](https://docs.docker.com/get-docker/).

Or Or you can install it by executing this command if you're on Debian or Ubuntu distribution:

```bash
sudo apt install docker.io
```
Well done. It's your turn to play now...

<a name="how_run_the_utility_on_machine_host"></a>
## How run the utility on machine host ?

To run this utility, you can execute this commands:

```bash
go build go_weather_go.go

./go_weather_go -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
or this single one:

```bash
go run go_weather_go.go -city=<wished_city_name> -apiKey=<OpenWeatherMap_API_key>
```
<a name="how_run_the_utility_with_docker"></a>
## How run the utility with Docker

To run this utility with Docker, you can build the Docker image from Dockerfile or you can pull the official Docker image on Docker Hub.

If you choose the last option, you can get the offical Docker image is [here](https://hub.docker.com/r/wicken/go_weather_go).

Now if you choose to build the Docker image from Dockerfile, follow these instructions in the order:

```bash
git clone https://gitlab.imerir.com/eric.ghoubiguian/go_weather_go # Clone the project from gitlab

cd go_weather_go # Change current folder for the project's one

docker image build -t go_weather_go . # Build the Docker image using Dockerfile
```
Now you are ready...

You can run the utility in Docker executing this command:

```bash
docker container run -it go_weather_go -city=<wished_city> -apiKey=<wished_api_key> -tempScale=<wished_temperature_scale>
```
Now let's go...
