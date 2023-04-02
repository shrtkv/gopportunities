.PHONY: default run build test docs clean

# Variables
APP_NAME=gopportunities

# Tasks

default: rdocs
run:
	@go run main.go
rdocs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -rf $(APP_NAME)
	@rm -rf ./docs
