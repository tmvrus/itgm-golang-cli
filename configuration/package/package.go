package main

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
)

type (
	FlagOptions struct {
		ApiKey    string `long:"api-key" description:"API access key" required:"true"`
		ApiSecret string `long:"api-secret" description:"API secret token" required:"true"`
		URL       string `long:"url" description:"Remote API endpoint" required:"true"`
	}

	EnvOptions struct {
		ApiKey    string        `env:"API_KEY,required"`
		ApiSecret string        `env:"API_SECRET,required"`
		URL       string        `env:"URL,required"`
		Timeout   time.Duration `env:"TIMEOUT" envDefault:"100ms"`
	}
)

func main() {

	cmdOps := FlagOptions{}
	argParser := flags.NewParser(&cmdOps, flags.None)
	if _, err := argParser.Parse(); err != nil {
		fmt.Printf("failed to Parse arguments: %s\n", err.Error())
		argParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	fmt.Println("Command line arguments")
	fmt.Printf("Key: %s\nSecret: %s\nURL: %s\n\n", cmdOps.ApiKey, cmdOps.ApiSecret, cmdOps.URL)

	envOps := EnvOptions{}
	if err := env.Parse(&envOps); err != nil {
		fmt.Printf("failed to Parse environment variables: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Environment variables arguments")
	fmt.Printf("Key: %s\nSecret: %s\nURL: %s\nTimeout: %s\n\n", envOps.ApiKey, envOps.ApiSecret, envOps.URL, envOps.Timeout)

	if err := godotenv.Overload(); err != nil { // godotenv.Load
		fmt.Printf("failed to Load .env file variables: %s", err.Error())
		os.Exit(1)
	}

	dotEnvOps := EnvOptions{}
	if err := env.Parse(&dotEnvOps); err != nil {
		fmt.Printf("failed to Parse environment variables: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("Config file arguments")
	fmt.Printf("Key: %s\nSecret: %s\nURL: %s\nTimeout: %s", dotEnvOps.ApiKey, dotEnvOps.ApiSecret, dotEnvOps.URL, dotEnvOps.Timeout)
}
