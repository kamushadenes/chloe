ifdef TARGET
	target := --target ${TARGET}
endif

.PHONY: all generate build clean

all: build

generate:
	go generate ./...

build: generate
	CGO_ENABLED=1 go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

clean:
	rm ./cmd/chloe/chloe

run: build
	cd cmd/chloe && ./chloe
