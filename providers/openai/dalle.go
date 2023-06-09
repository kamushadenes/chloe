package openai

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/models"
	putils "github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"os"
)

// cloneImageHeaders clones specified headers from a http.Response to an io.Writer.
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
func writeImage(writer structs.ChloeWriter, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	cloneImageHeaders(resp, writer)

	putils.WriteStatusCode(writer, resp.StatusCode)

	var buf bytes.Buffer

	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return err
	}

	return writer.WriteObject(&structs.ResponseObject{
		Name:   url,
		Data:   buf.Bytes(),
		Type:   structs.Image,
		Result: true,
	})
}

// getImageSize returns the appropriate image size for the request.
func getImageSize(request structs.ImageRequest) string {
	if request.GetSize() != "" {
		return request.GetSize()
	}

	return config.OpenAI.DefaultSize.ImageGeneration
}

// newImageRequest creates a new openai.ImageRequest for image generation.
func newImageRequest(request *structs.GenerationRequest) openai.ImageRequest {
	return openai.ImageRequest{
		Prompt: request.Prompt,
		N:      request.Count,
		Size:   request.Size,
	}
}

// createImageWithTimeout attempts to create an ImageResponse with a timeout.
func createImageWithTimeout(ctx context.Context, req openai.ImageRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeouts.WaitTimeout(ctx, config.Timeouts.ImageGeneration, func(ch chan interface{}, errCh chan error) {
		response, err := openAIClient.CreateImage(ctx, req)
		if err != nil {
			logger.Error().
				Err(err).
				Msg("error generating image")
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
	cost.AddCategoryCost(string(config.ImageGeneration),
		models.GetModel(fmt.Sprintf("dall-e-%s", request.Size)).UsageCost.Price*
			float64(len(response.Data)))

	channels.StartAndWait(request)

	for k := range response.Data {
		if err := writeImage(request.Writer, response.Data[k].URL); err != nil {
			return err
		}
	}

	return nil
}

// Generate generates an image based on a text prompt using the OpenAI API.
func Generate(request *structs.GenerationRequest) error {
	request.Size = getImageSize(request)

	logger := structs.LoggerFromRequest(request)

	logger.Info().
		Float64("estimatedCost",
			models.GetModel(fmt.Sprintf("dall-e-%s", request.Size)).UsageCost.Price*float64(request.Count)).
		Msg("generating image")

	req := newImageRequest(request)

	response, err := createImageWithTimeout(request.GetContext(), req)
	if err != nil {
		return channels.NotifyError(request, errors.ErrGenerationFailed, err)
	}

	err = processSuccessfulImageRequest(request, response)
	if err != nil {
		err = errors.Wrap(errors.ErrGenerationFailed, err)
	}

	return channels.NotifyError(request, err)
}

// newImageEditRequest creates a new openai.ImageEditRequest for image editing.
func newImageEditRequest(request *structs.GenerationRequest) (openai.ImageEditRequest, error) {
	f, err := os.Open(request.ImagePath)
	if err != nil {
		return openai.ImageEditRequest{}, err
	}

	return openai.ImageEditRequest{
		Prompt: request.Prompt,
		N:      request.Count,
		Size:   request.Size,
		Image:  f,
	}, nil
}

// createImageEditWithTimeout attempts to create an ImageResponse with a timeout for image editing.
func createImageEditWithTimeout(ctx context.Context, req openai.ImageEditRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeouts.WaitTimeout(ctx, config.Timeouts.ImageEdit, func(ch chan interface{}, errCh chan error) {
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

func processSuccessfulImageEditRequest(request *structs.GenerationRequest, response openai.ImageResponse) error {
	cost.AddCategoryCost(string(config.ImageEdit),
		models.GetModel(fmt.Sprintf("dall-e-%s", request.Size)).UsageCost.Price*
			float64(len(response.Data)))

	channels.StartAndWait(request)

	for k := range response.Data {
		if err := writeImage(request.Writer, response.Data[k].URL); err != nil {
			return err
		}
	}

	return nil
}

// Edits creates a new version of an image based on a text prompt using the OpenAI API.
func Edits(request *structs.GenerationRequest) error {
	request.Size = getImageSize(request)

	logger := structs.LoggerFromRequest(request)

	logger.Info().Msg("generating image edits")

	req, err := newImageEditRequest(request)
	if err != nil {
		return channels.NotifyError(request, errors.ErrGenerationFailed, err)
	}

	response, err := createImageEditWithTimeout(request.GetContext(), req)
	if err != nil {
		return channels.NotifyError(request, errors.ErrGenerationFailed, err)
	}

	err = processSuccessfulImageEditRequest(request, response)
	if err != nil {
		err = errors.Wrap(errors.ErrGenerationFailed, err)
	}

	return channels.NotifyError(request, err)
}

// newImageVariationRequest creates a new openai.ImageVariRequest for image variations.
func newImageVariationRequest(request *structs.VariationRequest) (openai.ImageVariRequest, error) {
	f, err := os.Open(request.ImagePath)
	if err != nil {
		return openai.ImageVariRequest{}, err
	}

	return openai.ImageVariRequest{
		Image: f,
		N:     request.Count,
		Size:  config.OpenAI.DefaultSize.ImageVariation,
	}, nil
}

// createImageVariationWithTimeout attempts to create an ImageResponse with a timeout for image variations.
func createImageVariationWithTimeout(ctx context.Context, req openai.ImageVariRequest) (openai.ImageResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeouts.WaitTimeout(ctx, config.Timeouts.ImageVariation, func(ch chan interface{}, errCh chan error) {
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
	cost.AddCategoryCost(string(config.ImageVariation),
		models.GetModel(fmt.Sprintf("dall-e-%s", request.Size)).UsageCost.Price*
			float64(len(response.Data)))

	channels.StartAndWait(request)

	for k := range response.Data {
		if err := writeImage(request.Writer, response.Data[k].URL); err != nil {
			return err
		}
	}

	return nil
}

// Variations generates variations of an input image using the OpenAI API.
func Variations(request *structs.VariationRequest) error {
	request.Size = getImageSize(request)

	logger := structs.LoggerFromRequest(request)

	logger.Info().Msg("generating image variations")

	req, err := newImageVariationRequest(request)
	if err != nil {
		return channels.NotifyError(request, errors.ErrGenerationFailed, err)
	}

	response, err := createImageVariationWithTimeout(request.GetContext(), req)
	if err != nil {
		return channels.NotifyError(request, errors.ErrGenerationFailed, err)
	}

	err = processSuccessfulImageVariationRequest(request, response)
	if err != nil {
		err = errors.Wrap(errors.ErrGenerationFailed, err)
	}

	return channels.NotifyError(request, err)
}
