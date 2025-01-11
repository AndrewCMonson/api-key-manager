#!/bin/bash

# Detect Operating System
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH="amd64" # Default architecture (adjust if supporting more)

# Set binary name and download URL
BINARY="oscarcli-$OS"
URL="https://github.com/AndrewCMonson/repo/releases/latest/download/$BINARY"

# Adjust for Windows
if [[ "$OS" == *"mingw"* || "$OS" == *"msys"* || "$OS" == "windows" ]]; then
  OS="windows"
  URL="https://github.com/AndrewCMonson/repo/releases/latest/download/oscarcli-windows.exe"
  DESTINATION="$HOME/oscarcli.exe"

else
  DESTINATION="/usr/local/bin/oscarcli"
fi

echo "Detected OS: $OS"
echo "Downloading CLI for $OS/$ARCH from $URL..."

# Download the binary
if [[ "$OS" == "windows" ]]; then
  curl -L -o "$DESTINATION" "$URL"
else
  sudo curl -L -o "$DESTINATION" "$URL"
  sudo chmod +x "$DESTINATION"
fi

# Success message
if [[ "$OS" == "windows" ]]; then
  echo "Installation complete. Run 'oscarcli' from the Command Prompt or PowerShell."
else
  echo "Installation complete. Run 'oscarcli' to get started."
fi
