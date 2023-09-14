ifdef TARGET
	target := --target ${TARGET}
endif

.PHONY: all generate build clean whisper test run

all: build

generate:
	go generate ./...

build: generate whisper
	cd cmd/chloe && LIBRARY_PATH=../../workspace/models/audio_models/whisper.cpp C_INCLUDE_PATH=../../workspace/models/audio_models/whisper.cpp CGO_ENABLED=1 go build

run:
	cd cmd/chloe && ./chloe

whisper:
	cd workspace/models/audio_models/whisper.cpp/bindings/go && make whisper > /dev/null
	cd workspace/models/audio_models/whisper.cpp && make main > /dev/null
	cd workspace/models/audio_models/whisper.cpp && bash ./models/download-ggml-model.sh medium

test: generate whisper
	LIBRARY_PATH=workspace/models/audio_models/whisper.cpp C_INCLUDE_PATH=workspace/models/audio_models/whisper.cpp CGO_ENABLED=1 go test ./...

clean:
	rm ./cmd/chloe/chloe