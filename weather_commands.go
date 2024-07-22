package main

import (
    "fmt"
)

func commandWeather(input string)error{ //So do i just call weather api here instead then?
    fmt.Println("this is the weather command")
    fmt.Printf("The city you passed was %s\n",input)
    return nil
}
