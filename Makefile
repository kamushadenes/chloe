# Makefile

.PHONY: all build clean

all: build_tokenizer build

build_tokenizer:
	cd tokenizer && cargo -C tiktoken-cffi build --release -Z unstable-options --out-dir .

build:
	go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

clean:
	rm ./cmd/chloe/chloe