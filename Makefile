ifdef TARGET
	target := --target ${TARGET}
endif

.PHONY: all generate build clean whisper test run

all: generate whisper build

generate:
	go generate ./...

build:
	cd cmd/chloe && LIBRARY_PATH=../../workspace/models/audio_models/whisper.cpp C_INCLUDE_PATH=../../workspace/models/audio_models/whisper.cpp CGO_ENABLED=1 go build

run:
	cd cmd/chloe && ./chloe

whisper:
	cd workspace/models/audio_models/whisper.cpp && make main
	cd workspace/models/audio_models/whisper.cpp && bash ./models/download-ggml-model.sh medium
	cd workspace/models/audio_models/whisper.cpp/bindings/go && make whisper

test:
	LIBRARY_PATH=workspace/models/audio_models/whisper.cpp C_INCLUDE_PATH=workspace/models/audio_models/whisper.cpp CGO_ENABLED=1 go test -v ./...

clean:
	rm ./cmd/chloe/chloe