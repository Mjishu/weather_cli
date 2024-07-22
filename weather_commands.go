package main

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

func commandWeather(input string)error{ //So do i just call weather api here instead then?
    fmt.Printf("city you passed was: %s\n",input)
    apiCall,err := apiCall("current.json", input)
    if err != nil{
        return err
    }
    fmt.Printf("Location: %s , Local Time: %s, Country: %s,\n", 
        apiCall.Location.Name, apiCall.Location.Localtime,apiCall.Location.Country)
    return nil
}

type apiResponse struct{
    Location struct{
        Name string `json:"name"`
        Localtime string `json:"localtime"`
        Country string `json:"country"`
    }`json:"location`
    Current struct {
        TempC float64 `json:"temp_c"`
        TempF float64 `json:"temp_f"`
        Condition struct{
            Text string `json:"text"`
        }`json:"condition"`
        Humidity int `json:"humidity"`
        Uv float64 `json:"uv"`
    }`jsonL"current"`
}

func apiCall(path, location string)(*apiResponse, error){ 
    api_key := os.Getenv("API_KEY")
    weather_url := "http://api.weatherapi.com/v1/" + path +"?key=" + api_key + "&q=" + location
    
    resp,err := http.Get(weather_url)
    if err != nil{
        return nil,err 
    }
    defer resp.Body.Close()

    body,err := ioutil.ReadAll(resp.Body)
    if err!=nil{
        return nil,err 
    }


    var apiResp apiResponse
    err = json.Unmarshal(body,&apiResp)
    if err != nil{
        return nil, fmt.Errorf("error parsing json %w",err)
    }

    return &apiResp,nil
}
