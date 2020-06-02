package main
import (
    "net/http"
    "fmt"
    "io/ioutil"
    "os"
)

func main() {
    resp, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=<city_name>&appid=<api_key>")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    text, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Println(string(text))
}
