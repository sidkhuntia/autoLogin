package main

import (
	"log"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	USERNAME := os.Getenv("USERNAME")
	PASSWORD := os.Getenv("PASSWORD")
	URL := os.Getenv("URL")

	var magicToken, inputURL string

	c.OnHTML("form", func(e *colly.HTMLElement) {
		e.ForEach("input[type='hidden']", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("name") {
			case "magic":
				magicToken = el.Attr("value")
			case "4Tredir":
				inputURL = el.Attr("value")
			}
		})
	})

	c.Visit(URL)

	// fmt.Printf("%v\n", magicToken)
	// fmt.Printf("%v\n", inputURL)

	// TODO: Check if the login was successful
	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode == 200 {
			fmt.Println("Login successful!")
		} else {
			fmt.Println("Login failed.")
		}
		fmt.Println(r)
	})

	c.Post(URL, map[string]string{
		"username": USERNAME,
		"password": PASSWORD,
		"magic":    magicToken,
		"4Tredir":  inputURL,
	})
}
