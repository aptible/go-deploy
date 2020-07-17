#!/usr/bin/env bash

mkdir -p bin
download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/tags/v0.21.0 | jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
curl -o bin/swagger -L'#' "$download_url"
chmod +x bin/swagger
