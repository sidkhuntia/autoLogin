package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func login() {
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

	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(strings.ToLower(string(r.Body)), "authentication failed") {
			fmt.Println("Authentication failed! Please check your credentials.")
		} else {
			fmt.Println("Login successful!")
		}
	})

	c.Post(URL, map[string]string{
		"username": USERNAME,
		"password": PASSWORD,
		"magic":    magicToken,
		"4Tredir":  inputURL,
	})

}

func checkIfConnectedToNetwor() bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	isConnectedToNetwork := false

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			isConnectedToNetwork = ipNet.IP != nil
		}
	}

	return isConnectedToNetwork
}

func main() {
	if checkIfConnectedToNetwor() {
		login()
	} else {
		fmt.Println("Not connected to a network.")
	}
}
