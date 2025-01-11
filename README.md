# Oscar CLI

Oscar CLI is a command-line tool for managing AWS Secrets Manager secrets. It allows you to create, update, and retrieve secrets, as well as generate API keys.

## Installation

1. Clone the repository:
    ```sh
    go install github.com/AndrewCMonson/oscarcli
    ```
## Usage

- You must have configured AWS credentials on your machine. You can do this by running `aws configure` and following the prompts.

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

- `apikey`: Generate a new API key and update the `oscar-api` secret in AWS Secrets Manager. Key length must be a valid integer.
    ```sh
    oscarcli apikey <key-length>
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
    oscarcli apikey 32
    ```

## License

This project is licensed under the MIT License.