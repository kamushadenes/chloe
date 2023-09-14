package whispercpp

import (
	"fmt"
	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/go-audio/wav"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"os"
	"path/filepath"
	"time"
)

type ASROptionsWhisperCpp struct {
	language     string
	timeout      time.Duration
	model        *common.ASRModel
	loadedModel  whisper.Model
	audioSamples []float32
	outFormat    string
}

func NewASROptionsWhisperCpp() common.ASROptions {
	return &ASROptionsWhisperCpp{}
}

func (c ASROptionsWhisperCpp) GetRequest() interface{} {
	return c.loadedModel
}

func (c ASROptionsWhisperCpp) GetAudioFile() string {
	return ""
}

func (c ASROptionsWhisperCpp) WithAudioFile(audioFile string) common.ASROptions {
	f, err := os.Open(audioFile)
	if err != nil {
		return c
	}
	defer f.Close()

	dec := wav.NewDecoder(f)

	if buf, err := dec.FullPCMBuffer(); err != nil {
		fmt.Println(fmt.Errorf("error decoding audio: %w", err))
	} else if dec.SampleRate != whisper.SampleRate {
		fmt.Println(fmt.Errorf("unsupported sample rate: %d", dec.SampleRate))
	} else if dec.NumChans != 1 {
		fmt.Println(fmt.Errorf("unsupported number of channels: %d", dec.NumChans))
	} else {
		c.audioSamples = buf.AsFloat32Buffer().Data
	}

	return c
}

func (c ASROptionsWhisperCpp) WithModel(model string) common.ASROptions {
	p := filepath.Join(config.Misc.WorkspaceDir, "models", "audio_models", "whisper.cpp", "models", fmt.Sprintf("ggml-%s.bin", model))
	m, err := whisper.New(p)
	if err != nil {
		fmt.Println(fmt.Errorf("error loading model: %w", err))
		return c
	}

	switch model {
	case "base":
		c.model = Base
	case "tiny":
		c.model = Tiny
	case "small":
		c.model = Small
	case "medium":
		c.model = Medium
	case "large":
		c.model = Large
	}

	c.loadedModel = m
	return c
}

func (c ASROptionsWhisperCpp) WithPrompt(prompt string) common.ASROptions {
	return c
}

func (c ASROptionsWhisperCpp) WithTemperature(temperature float32) common.ASROptions {
	return c
}

func (c ASROptionsWhisperCpp) WithLanguage(language string) common.ASROptions {
	c.language = language
	return c
}

func (c ASROptionsWhisperCpp) WithTimeout(timeout time.Duration) common.ASROptions {
	c.timeout = timeout
	return c
}

func (c ASROptionsWhisperCpp) GetTimeout() time.Duration {
	return c.timeout
}

func (c ASROptionsWhisperCpp) WithOutputFormat(format string) common.ASROptions {
	c.outFormat = format
	return c
}
