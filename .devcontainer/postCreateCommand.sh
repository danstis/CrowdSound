#!/bin/bash
set -e

echo alias ll=\\'ls -alF\\' >> ~/.bash_aliases

# Install the go helpers for this app:
# - swag - For the generation of the swagger docs.
go install github.com/swaggo/swag/cmd/swag@latest


# Install NodeJS for sonarlint.
curl -sL https://deb.nodesource.com/setup_16.x | sudo -E bash -
sudo apt-get install -y nodejs
