package envread

import (
	// "encoding/json"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AndrewCMonson/oscarcli/secrets"
)

// LoadEnvFile reads a .env file from the specified file path and returns a map of environment variables.
// The .env file should have lines in the format KEY=VALUE. Lines that are empty or start with a '#' are ignored.
// If a line does not contain an '=' character, an error is returned.
//
// Parameters:
//   - filePath: The path to the .env file to be read.
//
// Returns:
//   - A map where the keys are environment variable names and the values are the corresponding values.
//   - An error if the file cannot be opened, read, or contains invalid lines.
func LoadEnvFile(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open .env file: %w", err)
	}
	defer file.Close()

	envVars := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			// Skip empty lines or comments
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line in .env file: %s", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"")
		envVars[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading .env file: %w", err)
	}

	return envVars, nil
}

// UpdateSecretsFromEnvFile updates secrets in a specified secret manager from a .env file.
// It loads environment variables from the given file and updates each key-value pair in the secret manager.
//
// Parameters:
//   - secretname: The name of the secret to update.
//   - region: The region where the secret is stored.
//   - filePath: The path to the .env file containing the environment variables.
//
// Returns:
//   - error: An error if the .env file cannot be loaded or if any secret update fails.
func UpdateSecretsFromEnvFile(secretname, region, filePath string) error {
	envVars, err := LoadEnvFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to load .env file %s: %w", filePath, err)
	}

	for key, value := range envVars {
		fmt.Printf("Updating secret: %s=%s\n", key, value)
		if err := secrets.AddOrUpdateExistingSecret(secretname, region, key, value); err != nil {
			return fmt.Errorf("failed to update secret for key %s: %w", key, err)
		}
	}

	return nil
}