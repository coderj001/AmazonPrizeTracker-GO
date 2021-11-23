package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
    var in_url string

    fmt.Println(" :Welcome To Amazon Prize Tracker: ")
    fmt.Print("Enter URL: ")
    fmt.Scanln(&in_url)

    urlcheck(in_url)
}

func urlcheck(in_url string) bool {
    URL, err :=url.Parse(in_url)
    if err != nil {
        log.Panic(err)
    }
    fmt.Println(URL.Host)
    return true
}
