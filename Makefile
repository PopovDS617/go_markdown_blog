BINARY_NAME=myapp
HTTP_PORT="9000"
LOCAL_BIN:=$(CURDIR)/bin


## build: Build binary
build:
	@echo "Building..."
	env CGO_ENABLED=0  go build -ldflags="-s -w" -o ./bin/${BINARY_NAME} ./cmd
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env  HTTP_PORT=${HTTP_PORT} ./bin/${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ./bin/${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

race:
	cd ./cmd && \
	export HTTP_PORT=${HTTP_PORT} && \
	go run -race .

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./bin/${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...

testrace:
	go test -race -v ./...


testcov:
	go test -coverprofile=coverage.out -v ./...

testout:
	go tool cover -html=coverage.out

## docker up and down
du:
	docker-compose up -d
dd:
	docker-compose down