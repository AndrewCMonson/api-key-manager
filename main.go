package main

import (
	"fmt"
	"os"

	envwrite "github.com/AndrewCMonson/api-key-manager/envWrite"
	secrets "github.com/AndrewCMonson/api-key-manager/secrets"
)

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: oscar-secrets <command> <args>")
		os.Exit(1)
	}

	// Grab the command
	command := os.Args[1]

	// Handle each command
	switch command {
	case "env":
		if len(os.Args) != 4 {
			fmt.Println("Usage: oscar-secrets env <secret-name> <region>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]
		if err := envwrite.WriteENVToFile(secretName, region); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(".env successfully created/updated")

	case "create":
		if len(os.Args) != 6 {
			fmt.Println("Usage: oscar-secrets create <secret-name> <region> <key> <value>")
			os.Exit(1)
		}
		secretName := os.Args[2]
		region := os.Args[3]
		key := os.Args[4]
		value := os.Args[5]
		secrets.PushNewToSecretsManager(secretName, region, key, value)

	// case "update":
	// 	if len(os.Args) != 6 {
	// 		fmt.Println("Usage: oscar-secrets update <secret-name> <region> <key> <value>")
	// 		os.Exit(1)
	// 	}
	// 	secretName := os.Args[2]
	// 	region := os.Args[3]
	// 	key := os.Args[4]
	// 	value := os.Args[5]
	// 	handleUpdate(secretName, region, key, value)

	// case "apiGen":
	// 	if len(os.Args) != 4 {
	// 		fmt.Println("Usage: oscar-secrets apiGen <secret-name> <region>")
	// 		os.Exit(1)
	// 	}
	// 	secretName := os.Args[2]
	// 	region := os.Args[3]
	// 	handleApiGen(secretName, region)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands: env, create, update, apiGen")
		os.Exit(1)
	}
}