package main 

import (
    "github.com/joho/godotenv"
    "os"
    "fmt"
    "net/http"
    "bufio"
    "strings"
    //"io/ioutil"
)


func loadEnv(){
    err := godotenv.Load()
    if err !=nil{
        panic(err)
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

    /*data,err := ioutil.ReadAll(resp.Body)
    if err!=nil{
        return nil,err
    }
    var apiResponse apiResponse
    err = json.Unmarshal(data,&apiResponse)
    if err != nil{
        return nil, fmt.Errorf("error parsing json %w",err)
    }*/
    return nil,nil
}

type cliCommands struct{
    name string
    description string
    callback func()error
    callbackString func(string)error
}

func startRepl(){
    loadEnv()
    scanner := bufio.NewScanner(os.Stdin)
    /*_,err := apiCall("current.json", "jacksonville")
    if err != nil{
        fmt.Println("error calling api")
    }*/
    for {
        fmt.Print("Weather App > ")
        scanner.Scan()
        input := scanner.Text()

        words := cleanInput(input)
        if len(words) < 1 {
            return
        }
        commandName := words[0]
        command,exists := get_commands()[commandName]

        if exists{
            var err error
            if command.callbackString != nil{
                inputForCommand := strings.Join(words[1:]," ")
                err = command.callbackString(inputForCommand)
            }else if command.callback !=nil{
                err = command.callback()
            }
            if err !=nil{
                fmt.Println(err)
            }
        }else{
            fmt.Println("Unknown Command; try help")
        }
    }
}

func cleanInput(input string)[]string{
    lowered := strings.ToLower(input)
    split := strings.Fields(lowered)
    return split
}

func get_commands()map[string]cliCommands{
    return map[string]cliCommands{
        "help":{
            name:"help",
            description:"Shows different commands",
            callback: commandHelp,
        },
        "exit":{
            name:"exit",
            description:"Exits the program",
            callback: commandExit,
        },
        "weather":{
            name:"weather",
            description: "Shows the weather in x location | weather new york",
            callbackString: commandWeather,
        },
    }
}
