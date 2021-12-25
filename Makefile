all: manager

# Run manager binary
manager: fmt vet build

# Run go fmt againt code
fmt:
	go fmt ./...

# Run go vet against code
vet:
	go vet ./....

# Run go build against code
build:
	go build -o bin/go-setup main.go

# Run go run against code
run:
	go run main.go

# Run go mod tidy and vendor
mod:
	go mod tidy
	go mod vendor