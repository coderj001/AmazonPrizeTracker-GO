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

// Product struct
type Product struct {
	link          string
	product_name  string
	product_price float64
}

func main() {
	var input_url string

	fmt.Println(" :Welcome To Amazon Prize Tracker: ")
	fmt.Print("Enter URL: ")
	fmt.Scanln(&input_url)

	short_url, err := urlcheck(input_url)
	if err != nil {
		log.Panic(err)
	}

	product, err := scrapeweb(short_url)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(product)
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

// scrape amazon.in product page
// and return - {url, product_name, product_price}
func scrapeweb(short_url string) (Product, error) {
	client := http.Client{Timeout: 30 * time.Second}

	response, err := client.Get(short_url)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return Product{}, errors.New("Connection error")
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Panic(err)
	}

	product_name := strings.Trim(doc.Find(".a-size-large .product-title-word-break").Text(), "\n")

	product_price := doc.Find("#priceblock_ourprice").Text()
	product_price = strings.ReplaceAll(product_price, ",", "")
	product_price = strings.ReplaceAll(product_price, "₹", "")

	product_price_float, _ := strconv.ParseFloat(product_price, 64)

	return Product{link: short_url, product_name: product_name, product_price: product_price_float}, nil
}
