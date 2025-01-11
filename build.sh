#!/bin/bash

# Output directory
OUTPUT_DIR="build"
PLATFORMS=("linux/amd64" "darwin/amd64" "windows/amd64")

# Create the output directory
mkdir -p $OUTPUT_DIR

# Loop through platforms and build
for PLATFORM in "${PLATFORMS[@]}"; do
  GOOS=${PLATFORM%/*}   # Extract OS (e.g., "linux")
  GOARCH=${PLATFORM#*/} # Extract ARCH (e.g., "amd64")

  # Set the output file name
  OUTPUT_NAME="oscarcli-$GOOS"
  if [ "$GOOS" == "windows" ]; then
    OUTPUT_NAME+=".exe"
  fi

  echo "Building for $GOOS/$GOARCH..."
  GOOS=$GOOS GOARCH=$GOARCH go build -o "$OUTPUT_DIR/$OUTPUT_NAME"

  if [ $? -ne 0 ]; then
    echo "Failed to build for $GOOS/$GOARCH"
    exit 1
  fi
done

echo "Builds completed. Check the $OUTPUT_DIR directory."
