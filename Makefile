dayOne:
	go build -buildmode=plugin -o bin/1.so day/1/plugin.go

build:
	go build -o bin/main main.go

run:
	go run main.go

all: dayOne

go: all build run