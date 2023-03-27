package openai

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/flags"
	putils "github.com/kamushadenes/chloe/providers/utils"
	utils2 "github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeout"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

// cloneImageHeaders clones specified headers from an http.Response to an io.Writer.
func cloneImageHeaders(resp *http.Response, writer io.Writer) {
	for _, key := range []string{
		"Content-Type",
		"Content-Lenght",
		"Content-Disposition",
		"Content-MD5",
		"ETag",
	} {
		putils.CloneHeader(resp, writer, key)
	}
}

// writeImage writes an image from a URL to an io.WriteCloser.
func writeImage(request structs.Request, writer io.WriteCloser, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	cloneImageHeaders(resp, writer)

	putils.WriteStatusCode(writer, resp.StatusCode)

	if _, err := io.Copy(writer, resp.Body); err != nil {
		return err
	}

	return nil
}

// getImageSize returns the appropriate image size for the request.
func getImageSize(request structs.ImageRequest) string {
	if request.GetSize() != "" {
		return request.GetSize()
	}

	switch request.(type) {
	case *structs.GenerationRequest:
		return config.OpenAI.DefaultSize.ImageGeneration
	case *structs.VariationRequest:
		return config.OpenAI.DefaultSize.ImageVariation
	}

	return ""
}

// newImageRequest creates a new openai.ImageRequest for image generation.
func newImageRequest(request *structs.GenerationRequest) openai.ImageRequest {
	request.Size = getImageSize(request)

	return openai.ImageRequest{
		Prompt: request.Prompt,
		N:      len(request.GetWriters()),
		Size:   request.Size,
	}
}

// createImageWithTimeout attempts to create an ImageResponse with a timeout.
func createImageWithTimeout(ctx context.Context, req openai.ImageRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeout.WaitTimeout(ctx, config.Timeouts.ImageGeneration, func(ch chan interface{}, errCh chan error) {
		response, err := openAIClient.CreateImage(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error generating image")
			errCh <- err
		}
		ch <- response
	})
	if err != nil {
		return openai.ImageResponse{}, err
	}

	return respi.(openai.ImageResponse), nil
}

// processSuccessfulImageRequest processes a successful image generation request.
func processSuccessfulImageRequest(request *structs.GenerationRequest, response openai.ImageResponse) error {
	utils2.StartAndWait(request)

	for k := range request.GetWriters() {
		if err := writeImage(request, request.Writers[k], response.Data[k].URL); err != nil {
			return err
		}
	}

	return nil
}

// Generate generates an image based on a text prompt using the OpenAI API.
func Generate(request *structs.GenerationRequest) error {
	logger := structs.LoggerFromRequest(request)

	if flags.InteractiveCLI {
		return utils2.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image")

	req := newImageRequest(request)

	response, err := createImageWithTimeout(request.GetContext(), req)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	return utils2.NotifyError(request, processSuccessfulImageRequest(request, response))
}

// newImageEditRequest creates a new openai.ImageEditRequest for image editing.
func newImageEditRequest(request *structs.GenerationRequest) (openai.ImageEditRequest, error) {
	request.Size = getImageSize(request)

	f, err := os.Open(request.ImagePath)
	if err != nil {
		return openai.ImageEditRequest{}, err
	}

	return openai.ImageEditRequest{
		Prompt: request.Prompt,
		N:      len(request.Writers),
		Size:   request.Size,
		Image:  f,
	}, nil
}

// createImageEditWithTimeout attempts to create an ImageResponse with a timeout for image editing.
func createImageEditWithTimeout(ctx context.Context, req openai.ImageEditRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeout.WaitTimeout(ctx, config.Timeouts.ImageEdit, func(ch chan interface{}, errCh chan error) {
		response, err := openAIClient.CreateEditImage(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error generating image edits")
			errCh <- err
		}
		ch <- response
	})
	if err != nil {
		return openai.ImageResponse{}, err
	}

	return respi.(openai.ImageResponse), nil
}

// Edits creates a new version of an image based on a text prompt using the OpenAI API.
func Edits(request *structs.GenerationRequest) error {
	logger := structs.LoggerFromRequest(request)

	if flags.InteractiveCLI {
		return utils2.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image edits")

	req, err := newImageEditRequest(request)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	response, err := createImageEditWithTimeout(request.GetContext(), req)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	return utils2.NotifyError(request, processSuccessfulImageRequest(request, response))
}

// newImageVariationRequest creates a new openai.ImageVariRequest for image variations.
func newImageVariationRequest(request *structs.VariationRequest) (openai.ImageVariRequest, error) {
	request.Size = getImageSize(request)

	f, err := os.Open(request.ImagePath)
	if err != nil {
		return openai.ImageVariRequest{}, err
	}

	return openai.ImageVariRequest{
		Image: f,
		N:     len(request.Writers),
		Size:  config.OpenAI.DefaultSize.ImageVariation,
	}, nil
}

// createImageVariationWithTimeout attempts to create an ImageResponse with a timeout for image variations.
func createImageVariationWithTimeout(ctx context.Context, req openai.ImageVariRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeout.WaitTimeout(ctx, config.Timeouts.ImageVariation, func(ch chan interface{}, errCh chan error) {
		response, err := openAIClient.CreateVariImage(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error generating image variations")
			errCh <- err
		}
		ch <- response
	})
	if err != nil {
		return openai.ImageResponse{}, err
	}

	return respi.(openai.ImageResponse), nil
}

// processSuccessfulImageVariationRequest processes a successful image variation request.
func processSuccessfulImageVariationRequest(request *structs.VariationRequest, response openai.ImageResponse) error {
	utils2.StartAndWait(request)

	for k := range request.Writers {
		if err := writeImage(request, request.Writers[k], response.Data[k].URL); err != nil {
			return err
		}
	}

	return nil
}

// Variations generates variations of an input image using the OpenAI API.
func Variations(request *structs.VariationRequest) error {
	logger := structs.LoggerFromRequest(request)

	if flags.InteractiveCLI {
		return utils2.NotifyError(request, fmt.Errorf("can't generate images in CLI mode"))
	}

	logger.Info().Msg("generating image variations")

	req, err := newImageVariationRequest(request)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	response, err := createImageVariationWithTimeout(request.GetContext(), req)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	return utils2.NotifyError(request, processSuccessfulImageVariationRequest(request, response))
}
