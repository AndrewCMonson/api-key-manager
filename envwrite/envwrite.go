package envwrite

import (
	"context"
	"encoding/json"

	"fmt"
	"os"

	awsconfig "github.com/AndrewCMonson/oscarcli/services/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type AWSSecret struct {
	Name  string
	Value string
}

// this function grabs an entire "Secret" from AWS secrets manager.
// A "secret" can either be just a plaintext value or contain key value pairs
// 
// This function is meant to be used with a secret that contains key value pairs
// if used with a plaintext, it will throw an error explaining the cli tool isn't compatible
func getSecretFromSM(secretName, region string) (AWSSecret, error) {
	cfg, err := awsconfig.GetAWSConfig(region)
	if err != nil {
		return AWSSecret{}, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := secretsmanager.NewFromConfig(cfg)
	result, err := client.GetSecretValue(context.TODO(), &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		return AWSSecret{}, fmt.Errorf("failed to get secret value %w", err)
	}

	secret := AWSSecret{
		Name: *result.Name,
		Value: *result.SecretString,
	}

	return secret, nil
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
