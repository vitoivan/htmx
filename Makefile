
all: watch

watch: 
	@air -c ./.air.toml

run: 
	@go run ./src/main.go

build: 
	@go build ./src/main.go
