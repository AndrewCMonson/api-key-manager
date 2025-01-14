package secrets

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	awsconfig "github.com/AndrewCMonson/oscarcli/services/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

type AWSSecretKeyValue map[string]string

func generateAPIKey(length int) (string, error) {
	if length <= 31 {
		return "", errors.New("recommend a length of at least 32 for api key generation")
	}

	rawLength := (length * 6 + 7) / 8
	bytes := make([]byte, rawLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	key := base64.URLEncoding.EncodeToString(bytes)
	return key[:length], nil
}

func CreateNewAWSSecret(secretName, region, key, value string) error {
	cfg, err := awsconfig.GetAWSConfig(region)
	if err != nil {
		return fmt.Errorf("failed to load AWS config: %w", err)
	}

	secretValue := AWSSecretKeyValue{
		key: value,
	}

	client := secretsmanager.NewFromConfig(cfg)
	secretJson, err := json.Marshal(secretValue)
	if err != nil {
		return fmt.Errorf("error converting secret to json: %v", err)
	}

	_, err = client.CreateSecret(context.TODO(), &secretsmanager.CreateSecretInput{
		Name: 				aws.String(secretName),
		SecretString: aws.String(string(secretJson)),
	})
	if err != nil {
		return fmt.Errorf("failed to create secret %w", err)
	}

	return nil
}

// AddOrUpdateExistingSecret adds a new key:value pair to an existing secret
// or updates an existing key:value pair if the key already exists.
func AddOrUpdateExistingSecret(secretName, region, key, value string) error {
	cfg, err := awsconfig.GetAWSConfig(region)
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

func HandleAPIGen(action, secretname, region, key string, length int) (actionreturn, apikey string, err error) {
	apiKey, err := generateAPIKey(length)
	if err != nil {
		return "", "", fmt.Errorf("error generating API key: %v", err)
	}

	switch action {
	case "update":
		if err := AddOrUpdateExistingSecret(secretname, region, key, apiKey); err != nil {
			return "", "", fmt.Errorf("error: %v", err)
		}

		return action, apiKey, nil
	case "create":
		if err := CreateNewAWSSecret(secretname, region, key, apiKey); err != nil {
			return "", "", fmt.Errorf("error: %v", err)
		}

		return action, apiKey, nil

	default:
		return "", "", fmt.Errorf(`action must be either "update" or "create"`)
	}
}