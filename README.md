# Oscar CLI

Oscar CLI is a command-line tool for managing AWS Secrets Manager secrets. It allows you to create, update, and retrieve secrets, as well as generate API keys.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/AndrewCMonson/api-key-manager.git
    cd api-key-manager
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

### Commands

- `env`: Retrieve a secret from AWS Secrets Manager and write it to a [.env](http://_vscodecontentref_/1) file.
    ```sh
    oscarcli env <secret-name> <region>
    ```

- `create`: Create a new secret in AWS Secrets Manager.
    ```sh
    oscarcli create <secret-name> <region> <key> <value>
    ```

- `update`: Update an existing secret in AWS Secrets Manager.
    ```sh
    oscarcli update <secret-name> <region> <key> <value>
    ```

- `apiKeyGen`: Generate a new API key and update the `oscar-api` secret in AWS Secrets Manager.
    ```sh
    oscarcli apiKeyGen
    ```

### Examples

- Create a new secret:
    ```sh
    oscarcli create my-secret us-east-1 my-key my-value
    ```

- Update an existing secret:
    ```sh
    oscarcli update my-secret us-east-1 my-key new-value
    ```

- Retrieve a secret and write it to a [.env](http://_vscodecontentref_/2) file:
    ```sh
    oscarcli env my-secret us-east-1
    ```

- Generate a new API key:
    ```sh
    oscarcli apikey
    ```

## License

This project is licensed under the MIT License.