package openai

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

func writeImage(ctx context.Context, request structs.Request, writer io.WriteCloser, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return react.NotifyError(request, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	for _, key := range []string{
		"Content-Type",
		"Content-Lenght",
		"Content-Disposition",
		"Content-MD5",
		"ETag",
	} {
		cloneHeader(resp, writer, key)
	}

	writeStatusCode(writer, resp.StatusCode)

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return react.NotifyError(request, err)
	}

	if err := writer.Close(); err != nil {
		return react.NotifyError(request, err)
	}

	return nil
}

func Generate(ctx context.Context, request *structs.GenerationRequest) error {
	logger := zerolog.Ctx(ctx)

	if flags.CLI {
		return react.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image")

	if request.Size == "" {
		request.Size = config.OpenAI.DefaultSize[config.ImageTypeGeneration]
	}

	req := openai.ImageRequest{
		Prompt: request.Prompt,
		N:      len(request.Writers),
		Size:   request.Size,
	}

	response, err := openAIClient.CreateImage(ctx, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	react.StartAndWait(request)

	for k := range request.Writers {
		if err := writeImage(ctx, request, request.Writers[k], response.Data[k].URL); err != nil {
			return react.NotifyError(request, err)
		}
	}

	return react.NotifyError(request, nil)
}

func Edits(ctx context.Context, request *structs.GenerationRequest) error {
	logger := zerolog.Ctx(ctx)

	if flags.CLI {
		return react.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image edits")

	f, err := os.Open(request.ImagePath)
	if err != nil {
		return react.NotifyError(request, err)
	}

	if request.Size == "" {
		request.Size = config.OpenAI.DefaultSize[config.ImageTypeEdit]
	}

	req := openai.ImageEditRequest{
		Prompt: request.Prompt,
		N:      len(request.Writers),
		Size:   request.Size,
		Image:  f,
	}

	response, err := openAIClient.CreateEditImage(ctx, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	react.StartAndWait(request)

	for k := range request.Writers {
		if err := writeImage(ctx, request, request.Writers[k], response.Data[k].URL); err != nil {
			return react.NotifyError(request, err)
		}
	}

	return react.NotifyError(request, nil)
}

func Variations(ctx context.Context, request *structs.VariationRequest) error {
	logger := zerolog.Ctx(ctx)

	if flags.CLI {
		return react.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image variations")

	f, err := os.Open(request.ImagePath)
	if err != nil {
		return react.NotifyError(request, err)
	}
	defer f.Close()

	req := openai.ImageVariRequest{
		Image: f,
		N:     len(request.Writers),
		Size:  config.OpenAI.DefaultSize[config.ImageTypeVariation],
	}

	response, err := openAIClient.CreateVariImage(ctx, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	react.StartAndWait(request)

	for k := range request.Writers {
		resp, err := http.Get(response.Data[k].URL)
		if err != nil {
			return react.NotifyError(request, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status: %s", resp.Status)
		}

		if _, err = io.Copy(request.Writers[k], resp.Body); err != nil {
			return react.NotifyError(request, err)
		}

		if err := request.Writers[k].Close(); err != nil {
			return react.NotifyError(request, err)
		}
	}

	return react.NotifyError(request, nil)
}
