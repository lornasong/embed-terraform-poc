
PROJECT = $(shell basename $(CURDIR))

deps:
	go mod download

run:
	go build -o build/$(PROJECT) && ./build/$(PROJECT)
