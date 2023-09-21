ifdef TARGET
	target := --target ${TARGET}
endif

.PHONY: all generate build clean whisper test run

ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
LIBRARY_PATH := ${ROOT_DIR}/workspace/models/audio_models/whisper.cpp
C_INCLUDE_PATH := ${ROOT_DIR}/workspace/models/audio_models/whisper.cpp
CGO_ENABLED := 1

all: generate whisper build

generate:
	go generate ./...

build:
	cd cmd/chloe && C_INCLUDE_PATH=${C_INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} CGO_ENABLED=${CGO_ENABLED} go build

run:
	cd cmd/chloe && ./chloe

whisper:
	cd workspace/models/audio_models/whisper.cpp && make main
	cd workspace/models/audio_models/whisper.cpp && bash ./models/download-ggml-model.sh medium
	cd workspace/models/audio_models/whisper.cpp/bindings/go && make whisper

test:
	C_INCLUDE_PATH=${C_INCLUDE_PATH} LIBRARY_PATH=${LIBRARY_PATH} CGO_ENABLED=${CGO_ENABLED} go test -v ./...

clean:
	rm ./cmd/chloe/chloe