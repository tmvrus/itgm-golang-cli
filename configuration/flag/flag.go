package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	defaultFlagValue = ""
)

func main() {
	var apiKey, apiSecret, url string
	flag.StringVar(&apiKey, "api-key", defaultFlagValue, "API access key")
	flag.StringVar(&apiSecret, "api-secret", defaultFlagValue, "API secret")
	flag.StringVar(&url, "url", defaultFlagValue, "API URL endpoint")

	flag.Parse()

	if apiKey == defaultFlagValue || apiSecret == defaultFlagValue || url == defaultFlagValue {
		fmt.Println("not all arguments defined")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Printf("Key: %s\nSecret: %s\nURL: %s\n", apiKey, apiSecret, url)
}
