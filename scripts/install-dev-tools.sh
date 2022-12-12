#!/bin/sh -e

# Install development tools
mkdir -p tools
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1
go install github.com/patrickglass/temporalite/cmd/temporalite@latest

# Install redis server
curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
sudo apt-get update
sudo apt-get install -yq redis
sudo apt-get clean -y

# Install Temporal Server
# temporalite start --ephemeral --namespace default
