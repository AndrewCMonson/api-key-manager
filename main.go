package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndrewCMonson/oscarcli/envread"
	"github.com/AndrewCMonson/oscarcli/envwrite"
	"github.com/AndrewCMonson/oscarcli/secrets"
)

const version = "1.4.0"

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: oscarcli <command> <args>")
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
	case "env-get":
		if len(os.Args) != 4 {
			fmt.Println("Usage: oscarcli env-get <secret-name> <region>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]

		if err := envwrite.WriteENVToFile(secretName, region); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(".env successfully created/updated")
	case "env-set":
		if len(os.Args) != 5 {
			fmt.Println("Usage: oscarcli env-set <secret-name> <region> <env-file-path>")
			os.Exit(1)
		}

		secretName := os.Args[2]
		region := os.Args[3]
		filePath := os.Args[4]

		if err := envread.UpdateSecretsFromEnvFile(secretName, region, filePath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("AWS Secrets Manager successfully updated")
	case "env-create":
		if len(os.Args) != 5 {
			fmt.Println("Usage: oscarcli env-create <secret-name> <region> <env-file-path>")
			os.Exit(1)
		}

		secretName := os.Args[2]
		region := os.Args[3]
		filePath := os.Args[4]

		if err := envread.CreateAndWriteSecretsFromEnv(secretName, region, filePath); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Successfully created new AWS Secret with .env values named %s\n", secretName)

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
		if len(os.Args) != 7 {
			fmt.Println("Usage: oscarcli apikey <action: create|update> <secret-name> <region> <api-key-name> <key-length(int)>")
			os.Exit(1)
		}

		action := os.Args[2]
		secretname := os.Args[3]
		region := os.Args[4]
		secretkey := os.Args[5]
		lengthParam := os.Args[6]

		length, err := strconv.Atoi(lengthParam)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		a, k, err := secrets.HandleAPIGen(action, secretname, region, secretkey, length)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		switch a {
		case "create": 
			fmt.Printf("API key successfully created!\nAWS Secret Name: %s\n%s:%s\n", secretname, secretkey, k)
		 
		case "update":
			fmt.Printf("API key successfully updated!\nAWS Secret Name: %s\n%s:%s\n", secretname, secretkey, k)

		default:
			fmt.Printf("Unknown api key action\n")
			os.Exit(1)
		}
	
	
		default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: env-get, env-set, env-create, create, update, apikey")
		os.Exit(1)
	}
}