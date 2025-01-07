package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func generateAPIKey(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func pushToSecretsManager(secretName, apiKey, region string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))

	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)

	_, err = client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
		Name: 				aws.String(secretName),
		SecretString: aws.String(apiKey),
	})
	if err != nil {
		return fmt.Errorf("failed to create secret %w", err)
	}

	fmt.Printf("API key successfully stored in AWS secrets manager under name %s\n", secretName)
	return nil
}

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <secret-name> <region>", os.Args[0])
	}

	secretName := os.Args[1]
	region := os.Args[2]

	apiKey, err := generateAPIKey(32)
	if err != nil {
		log.Fatalf("error generating API key: %v", err)
	}

	fmt.Printf("Generated API key: %s\n", apiKey)

	if err := pushToSecretsManager(secretName, apiKey, region); err != nil {
		log.Fatalf("error p")
	}
}