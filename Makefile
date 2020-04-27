
PROJECT = $(shell basename $(CURDIR))

deps:
	go mod download

apply:
	go build -o build/$(PROJECT) && ./build/$(PROJECT)

destroy:
	go build -o build/$(PROJECT) && ./build/$(PROJECT) -destroy=true
