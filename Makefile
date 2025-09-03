.PHONY: swagger build start

all: swagger build

swagger:
	swag init -g cmd/main.go

build:
	go build -o dist/api cmd/main.go
