package apikeygen

import (
	// "bufio"
	"context"
	// "crypto/rand"
	// "encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

type AWSSecretKeyValue map[string]string

// func generateAPIKey(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}

// 	return hex.EncodeToString(bytes), nil
// }

func CreateNewAWSSecret(secretName, region, key, value string) error {
  cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	secretValue := AWSSecretKeyValue{
		key: value,
	}

	client := secretsmanager.NewFromConfig(cfg)
	secretJson, err := json.Marshal(secretValue)
	if err != nil {
		log.Fatalf("error converting secret to json: %v", err)
	}

	_, err = client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
		Name: 				aws.String(secretName),
		SecretString: aws.String(string(secretJson)),
	})
	if err != nil {
		return fmt.Errorf("failed to create secret %w", err)
	}

	fmt.Printf("Secret successfully stored in AWS secrets manager under name %s\n", secretName)
	return nil
}

// AddOrUpdateExistingSecret adds a new key:value pair to an existing secret
// or updates an existing key:value pair if the key already exists.
func AddOrUpdateExistingSecret(secretName, region, key, value string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)

	// Retrieve the current secret value
	getSecretResp, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		if _, ok := err.(*types.ResourceNotFoundException); ok {
			return fmt.Errorf("secret %s not found", secretName)
		}
		return fmt.Errorf("failed to retrieve secret: %w", err)
	}

	// Deserialize the current secret value
	var currentSecret map[string]string
	if err := json.Unmarshal([]byte(aws.ToString(getSecretResp.SecretString)), &currentSecret); err != nil {
		return fmt.Errorf("failed to parse current secret: %w", err)
	}

	// Add or update the key-value pair
	currentSecret[key] = value

	// Serialize the updated secret value
	updatedSecretJson, err := json.Marshal(currentSecret)
	if err != nil {
		return fmt.Errorf("failed to serialize updated secret: %w", err)
	}

	// Update the secret in AWS Secrets Manager
	_, err = client.UpdateSecret(context.TODO(), &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(secretName),
		SecretString: aws.String(string(updatedSecretJson)),
	})
	if err != nil {
		return fmt.Errorf("failed to update secret: %w", err)
	}

	return nil
}
// func GenerateAndPushSecret() {
// 	if len(os.Args) < 5 {
// 		log.Fatalf("Usage: %s <secret-name> <region>", os.Args[0])
// 	}

// 	var storedSecret AWSSecretKeyValue

// 	secretName := os.Args[1]
// 	region := os.Args[2]
// 	storedSecret.Key = os.Args[3]
// 	storedSecret.Value = os.Args[4]

// 	// apiKey, err := generateAPIKey(32)
// 	// if err != nil {
// 	// 	log.Fatalf("error generating API key: %v", err)
// 	// }

// 	// fmt.Printf("Generated API key: %s\n", apiKey)

// 	if err := PushNewToSecretsManager(secretName, region, storedSecret); err != nil {
// 		log.Fatalf("error pushing secret to secrets manager: %v", err)
// 	}

// 	fmt.Println("Secret Generated and Pushed to AWS Secrets Manager")
// }