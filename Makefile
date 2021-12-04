DAY?=all

days: one two three

one:
	go build -buildmode=plugin -o bin/1.so days/1/one.go
two:
	go build -buildmode=plugin -o bin/2.so days/2/two.go
three:
	go build -buildmode=plugin -o bin/3.so days/3/three.go

build:
	go build -o bin/main main.go

all: days build 

run:
	go run main.go -day=$(DAY)

clean:
	rm -f bin/*.so
	rm -f bin/main