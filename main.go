package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	title := doc.Find("span .a-size-large .product-title-word-break")
	fmt.Println(title.Text())
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
	defer resp.Body.Close()

	return resp
}
