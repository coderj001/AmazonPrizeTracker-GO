package main

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	var in_url string

	fmt.Println(" :Welcome To Amazon Prize Tracker: ")
	fmt.Print("Enter URL: ")
	fmt.Scanln(&in_url)

	rurl, err := urlcheck(in_url)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(rurl)
}

func urlcheck(in_url string) (string, error) {
	var short_url string
	URL, err := url.Parse(in_url)
	if err != nil {
		log.Panic(err)
	}

	if URL.Host != "www.amazon.in" {
		return "", errors.New("Host of the url should be 'www.amazon.in'.")
	}
	if URL.Scheme != "https" {
		return "", errors.New("Should be https.")
	}
	str_path := strings.Split(URL.Path, "/")
	short_url = URL.Scheme + "://" + URL.Host + "/" + str_path[2] + "/" + str_path[3]
	return short_url, nil
}
