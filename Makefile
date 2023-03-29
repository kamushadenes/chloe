ifdef TARGET
	target := --target ${TARGET}
endif

.PHONY: all tokenizer generate build clean

all: tokenizer build

tokenizer:
	cd tokenizer && cargo -C tiktoken-cffi build ${target} --release -Z unstable-options --out-dir .

generate:
	go generate ./...

build: tokenizer generate
	CGO_ENABLED=1 go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

clean:
	rm ./cmd/chloe/chloe

run: build
	cd cmd/chloe && ./chloe