package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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

	response := scrapeweb(rurl)
	fmt.Println(response.Status)
	defer response.Body.Close()
	doc, _ := goquery.NewDocumentFromReader(response.Body)

	product_name := strings.Trim(doc.Find(".a-size-large .product-title-word-break").Text(), "\n")
	fmt.Println("Poduct Name: ", product_name)

	product_price := doc.Find("#priceblock_ourprice").Text()
	fmt.Println("Poduct Prize: ", product_price[2:])
    // TODO:  <24-11-21, coderj001> // Error not convert to float
	product_price_float, _ := strconv.ParseFloat(product_price[2:], 64)
	fmt.Println("Poduct Prize: ", product_price_float)
}

// convert bulk url into sutiable formant
// ex: http://www.amazon.com/dp/JA12AN120
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

func scrapeweb(short_url string) *http.Response {
	client := http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(short_url)
	if err != nil {
		log.Panic(err)
	}

	return resp
}
