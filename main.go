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

func login() error {
	c := colly.NewCollector()
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("failed to load environment variables: %w", err)
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

	err = c.Visit(URL)
	if err != nil {
		return fmt.Errorf("failed to visit URL: %w", err)
	}

	c.OnResponse(func(r *colly.Response) {
		if strings.Contains(strings.ToLower(string(r.Body)), "authentication failed") {
			fmt.Println("Authentication failed! Please check your credentials.")
		} else {
			fmt.Println("Login successful!")
		}
	})

	err = c.Post(URL, map[string]string{
		"username": USERNAME,
		"password": PASSWORD,
		"magic":    magicToken,
		"4Tredir":  inputURL,
	})
	if err != nil {
		return fmt.Errorf("failed to post login credentials: %w", err)
	}

	return nil
}

func checkIfConnectedToNetwork() bool {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("failed to get network interface addresses: %s", err)
	}

	isConnectedToNetwork := false

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			isConnectedToNetwork = ipNet.IP != nil
			if isConnectedToNetwork {
				break
			}
		}
	}

	return isConnectedToNetwork
}

func main() {
	if checkIfConnectedToNetwork() {
		err := login()
		if err != nil {
			log.Fatalf("failed to login: %s", err)
		}
	} else {
		fmt.Println("Not connected to a network.")
	}
}
