# Makefile

.PHONY: all tokenizer build clean

all: tokenizer build

tokenizer:
	cd tokenizer && cargo -C tiktoken-cffi build --release -Z unstable-options --out-dir .

build:
	go build -o ./cmd/chloe/chloe ./cmd/chloe/main.go

clean:
	rm ./cmd/chloe/chloe