.PHONY: all build clean

all: build

build: clean
	go build -o bin/member cmd/member/main.go

run:
	./bin/member

clean:
	rm -rf bin/member