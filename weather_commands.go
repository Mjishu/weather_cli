package main

import (
    "fmt"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strings"
)

func commandUv(input string)error{
    apiCall,err := apiCall("current.json",input)
    if err != nil{
        return err
    }
    fmt.Println()
    fmt.Println("           Location")
    fmt.Println(strings.Repeat("-",30))
    fmt.Printf("City: %s\nLocal Time: %s\nCountry: %s\n", 
        apiCall.Location.Name, apiCall.Location.Localtime,apiCall.Location.Country)
    fmt.Println()
    fmt.Println("           Heat and UV")
    fmt.Println(strings.Repeat("-",30))
    fmt.Printf("UV index: %.1f\nHeat Index:%.1f\n",apiCall.Current.Uv,apiCall.Current.Heatindex_f)

    return nil
}

func commandWeather(input string)error{ //So do i just call weather api here instead then?
    fmt.Println()
    apiCall,err := apiCall("current.json", input)
    if err != nil{
        return err
    }
    fmt.Println("           Location")
    fmt.Println(strings.Repeat("-",30))

    fmt.Printf("City: %s\nLocal Time: %s\nCountry: %s\n", 
        apiCall.Location.Name, apiCall.Location.Localtime,apiCall.Location.Country)

    fmt.Println()
    fmt.Println("           weather")
    fmt.Println(strings.Repeat("-",30))

    fmt.Printf("Temperature: %.1f F\nHumidity: %d%%\nCondition: %s\n",
        apiCall.Current.TempF,apiCall.Current.Humidity, apiCall.Current.Condition.Text)
    fmt.Println()
    fmt.Println(strings.Repeat("-", 30))

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
            Icon string `json:"icon"`
        }`json:"condition"`
        Humidity int `json:"humidity"`
        Uv float64 `json:"uv"`
        Heatindex_c float64 `json:"heatindex_c"`
        Heatindex_f float64 `json:"heatindex_f"`

    }`json:"current"`
    Error struct{
        Code int `json:"code"`
        Message string `json:"message"`
    }`json:"error"`
}

func apiCall(path, location string)(*apiResponse, error){ 
    api_key := os.Getenv("API_KEY")
    weather_url := "http://api.weatherapi.com/v1/" + path +"?key=" + api_key + "&q=" + location
    
    // check for error.code and return maybe the error.message?
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
