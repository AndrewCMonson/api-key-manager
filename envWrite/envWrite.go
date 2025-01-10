package envwrite

import (
	"context"
	"encoding/json"

	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AWSSecret struct {
	Name  string
	Value string
}

// this function grabs an entire "Secret" from AWS secrets manager.
// A "secret" and either be just a plaintext value or contain key value pairs
// 
// This function is meant to be used with a secret that contains key value pairs
// if used with a plaintext, it will throw an error explaining
func getSecretFromSM(secretName, region string) (secret AWSSecret, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return AWSSecret{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	var secrets AWSSecret
	result, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return AWSSecret{}, fmt.Errorf("failed to get secret value %w", err)
	}

	secrets.Name = *result.Name
	secrets.Value = *result.SecretString

	return secrets, nil
}

func WriteENVToFile(secretName, region string) error {
	secretValue, err := getSecretFromSM(secretName, region)
	if err != nil {
		return fmt.Errorf("error retrieving secret from Secrets Manager: %w", err)
	}

	var secretMap map[string]string
	if err := json.Unmarshal([]byte(secretValue.Value), &secretMap); err != nil {
		return fmt.Errorf("not compatible with plaintext secrets")
	}

	file, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("error creating .env file: %w", err)
	}
	defer file.Close()

	for key, value := range secretMap {
		if _, err := file.WriteString(fmt.Sprintf("%s=\"%s\"\n", key, value)); err != nil {
			return fmt.Errorf("error writing to .env file: %w", err)
		}
	}

	

	return nil
}
