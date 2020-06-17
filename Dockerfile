#Put the golang image as image's base
FROM golang

#
LABEL maintainer="ericghoubiguian@live.fr"

#Install the gjson project to manipulate the JSON response from HTTP request
RUN go get -u github.com/tidwall/gjson

#Copy all files in the newly created directory go_weather_go
COPY . /go_weather_go

#Change work directory for the go_weather_go project one
WORKDIR /go_weather_go

#Container instruction as entrypoint: 'go run go_weather_go'
ENTRYPOINT ["go", "run", "go_weather_go.go"]
