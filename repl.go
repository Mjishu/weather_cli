package main 

import (
    "github.com/joho/godotenv"
    "os"
    "fmt"
    "net/http"
    "io/ioutil"
)
func loadEnv(){
    err := godotenv.Load()
    if err !=nil{
        panic(err)
    }
}

}

func apiCall(path, location string)([]string, error){
    api_key := os.Getenv("API_KEY")
    weather_url := "http://api.weatherapi.com/v1" + path +"?key=" + api_key + "&q=" + location
    resp,err := http.Get(weather_url)
    if err != nil{
        return nil,err
    }
    defer resp.Body.Close()

    data,err := ioutil.ReadAll(resp.Body)
    if err!=nil{
        return nil,err
    }
    /*var apiResponse apiResponse
    err = json.Unmarshal(data,&apiResponse)
    if err != nil{
        return nil, fmt.Errorf("error parsing json %w",err)
    }*/
}

func startRepl(){
    loadEnv()
    data,err := apiCall("current.json", "jacksonville")
    if err != nil{
        fmt.Println("error calling api")
    }
    fmt.Println(data)
}
