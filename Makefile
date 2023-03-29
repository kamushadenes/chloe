# Makefile

.PHONY: all tokenizer build clean

all: tokenizer build

tokenizer:
	cd tokenizer && cargo -C tiktoken-cffi build --release -Z unstable-options --out-dir .

build: tokenizer
	CGO_ENABLED=1 go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

clean:
	rm ./cmd/chloe/chloe

run: build
	cd cmd/chloe && ./chloe