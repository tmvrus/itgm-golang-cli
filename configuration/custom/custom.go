package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	envConfigKey = "CONFIG_FILE_PATH"
)

type (
	Config struct {
		ApiKey    string `json:"api_key"`
		ApiSecret string `json:"api_secret"`
		URL       string `json:"url"`
	}
)

func main() {
	fmt.Printf("%T\n", os.Args)
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	fileName, present := os.LookupEnv(envConfigKey)
	if !present {
		fmt.Println("environment variable CONFIG_FILE_PATH not provided")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("failed to ReadFile %q, error: %q", fileName, err.Error())
		os.Exit(1)
	}

	cfg := Config{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		fmt.Printf("failed to Unmarshal data, error: %q", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Key: %s\nSecret: %s\nURL: %s\n", cfg.ApiKey, cfg.ApiSecret, cfg.URL)
}
