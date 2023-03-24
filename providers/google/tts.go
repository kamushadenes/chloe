package google

import (
	"bytes"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"io"
)

func TTS(ctx context.Context, request *structs.TTSRequest) error {
	logger := zerolog.Ctx(ctx)

	if flags.InteractiveCLI {
		return fmt.Errorf("can't generate audio in CLI mode")
	}

	logger.Info().Msg("converting text to speech")

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: request.Content},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: config.GCP.TTSLanguageCode,
			Name:         config.GCP.TTSVoiceName,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: config.GCP.TTSEncoding,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return err
	}

	var contentType string

	switch config.GCP.TTSEncoding {
	case texttospeechpb.AudioEncoding_MP3:
		contentType = "audio/mpeg"
	case texttospeechpb.AudioEncoding_OGG_OPUS:
		contentType = "audio/ogg"
	case texttospeechpb.AudioEncoding_LINEAR16:
		contentType = "audio/wav"
	case texttospeechpb.AudioEncoding_MULAW:
		contentType = "audio/mulaw"
	case texttospeechpb.AudioEncoding_ALAW:
		contentType = "audio/alaw"
	}

	for k := range request.Writers {
		utils.WriteHeader(request.Writers[k], "Content-Type", contentType)
		utils.WriteHeader(request.Writers[k], "Content-Length", fmt.Sprintf("%d", len(resp.AudioContent)))

		if _, err := io.Copy(request.Writers[k], bytes.NewReader(resp.AudioContent)); err != nil {
			return err
		}
		if err := request.Writers[k].Close(); err != nil {
			return err
		}
	}

	return nil
}
