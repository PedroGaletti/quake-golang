install:
	@echo "Downloading dependecies..."
	@go get
	@echo "Validating dependecies..."
	@go mod tidy

build:
	@echo "Building project..."
	@go build
	@echo "Build completed successfully."

run:
	@echo "Running application..."
	@go run main.go

clean:
	@echo "Cleaning up project..."
	@rm -rf ./go.sum
	@rm -rf ./quake
	@echo "Project cleaned successfully."