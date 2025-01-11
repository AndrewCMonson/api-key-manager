package main

import (
	"fmt"
	"os"
	"strconv"

	envwrite "github.com/AndrewCMonson/oscarcli/envWrite"
	"github.com/AndrewCMonson/oscarcli/secrets"
)

const version = "1.2.1"

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: oscar-secrets <command> <args>")
		os.Exit(1)
	}

	// Grab the command
	command := os.Args[1]

	// Handle version flag
	if command == "--version" || command == "-v" {
		fmt.Printf("oscarcli version %s\n", version)
		os.Exit(0)
	}

	// Handle each command
	switch command {
	case "env":
		if len(os.Args) != 4 {
			fmt.Println("Usage: oscarcli env <secret-name> <region>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]
		if err := envwrite.WriteENVToFile(secretName, region); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(".env successfully created/updated")

	case "create":
		if len(os.Args) != 6 {
			fmt.Println("Usage: oscarcli create <secret-name> <region> <key> <value>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]
		key := os.Args[4]
		value := os.Args[5]
		if err := secrets.CreateNewAWSSecret(secretName, region, key, value); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Secret successfully created in AWS Secrets Manager under name %s\n", secretName)

	case "update":
		if len(os.Args) != 6 {
			fmt.Println("Usage: oscarcli update <secret-name> <region> <key> <value>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]
		key := os.Args[4]
		value := os.Args[5]
		if err := secrets.AddOrUpdateExistingSecret(secretName, region, key, value); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Secret successfully updated in AWS Secrets Manager under name %s\n", secretName)

	case "apikey":
		if len(os.Args) != 3 {
			fmt.Println("Usage: oscarcli apiKeyGen <key-length(int)>")
			os.Exit(1)
		}

		lengthParam := os.Args[2]

		length, err := strconv.Atoi(lengthParam)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		if err := secrets.HandleAPIGen(length); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Oscar API key successfully updated!")

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: env, create, update, apikey")
		os.Exit(1)
	}
}