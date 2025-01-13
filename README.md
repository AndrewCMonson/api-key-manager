# OscarCLI

OscarCLI is a command-line tool for managing AWS Secrets Manager secrets. It allows you to create, update, and retrieve secrets, as well as generate API keys. It is also used to write secrets to a [.env](http://_vscodecontentref_/0) file for use in local development.

## Installation

### Using Go

1. If Go is already installed, run the following command to install OscarCLI:
    ```sh
    go install github.com/AndrewCMonson/oscarcli@latest
    ```

### Manual Installation

1. If Go is not installed, go to the [Releases Page](https://github.com/username/repo/releases).
   - Download the binary for your operating system:
     - **Linux**: `oscarcli-linux`
     - **MacOS**: `oscarcli-darwin`
     - **Windows**: `oscarcli-windows.exe`
   
2. Move the binary to a directory in your `PATH` and make it executable:
   - **Linux/MacOS**:
     ```bash
     mv oscarcli-linux /usr/local/bin/oscarcli
     chmod +x /usr/local/bin/oscarcli
     ```
   - **Windows**:
     Rename `oscarcli-windows.exe` to `oscarcli.exe` and move it to `C:\Windows\System32`.

3. Verify the installation by running `oscarcli --version or oscarcli -v`.

## Usage

- You must have configured AWS credentials on your machine. You can do this by running `aws configure` and following the prompts.
- As of version 1.2.4, OscarCLI uses the credentials for AWS of the user running the command. This means that the user must have the necessary permissions to create, update, and retrieve secrets in AWS Secrets Manager.
- Using the `apikey` command will update the user's `oscar-api` secret in AWS Secrets Manager. This will not update the master API key used by the Oscar API. Future versions of OscarCLI will include the ability to update the master API key, locked behind credential verification.
- When using the `env-get` command, the `.env` file will be created in the current working directory. If the file already exists, it will be overwritten.
- When using the `env-set` command, the `.env` file will be read from the specified path. If the file does not exist, an error will be thrown.

### Commands

- `env-get`: Retrieve a secret from AWS Secrets Manager and write it to a [.env](http://_vscodecontentref_/1) file.
    ```sh
    oscarcli env <secret-name> <region>
    ```
- `env-set`: Write a .env file to AWS Secrets Manager.
    ```sh
    oscarcli env-set <secret-name> <region> <env-file-path>
    ```
- `env-create`: Create a new AWS Secret and populate it with a .env file's contents.
    ```sh
    oscarcli env-create <secret-name> <region> <env-file-path>
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
    oscarcli env-get my-secret us-east-1
    ```

- Write a .env file to AWS Secrets Manager:
    ```sh
    oscarcli env-set my-secret us-east-1 /path/to/.env
    ```

- Create a new secret and populate it with a .env file's contents:
    ```sh
    oscarcli env-create my-secret us-east-1 /path/to/.env
    ```

- Generate a new API key:
    ```sh
    oscarcli apikey 32
    ```

## License

This project is licensed under the MIT License.
