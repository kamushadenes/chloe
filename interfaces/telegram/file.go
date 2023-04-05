package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"io"
	"net/http"
	"os"
	"path"
)

func getFile(api *tgbotapi.BotAPI, fileID string) (tgbotapi.File, error) {
	return api.GetFile(tgbotapi.FileConfig{FileID: fileID})
}

func createRequest(url string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, url, nil)
}

func downloadFileData(api *tgbotapi.BotAPI, req *http.Request) ([]byte, error) {
	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return io.ReadAll(resp.Body)
}

func saveFile(filePath string, data []byte) error {
	if err := os.MkdirAll(path.Dir(filePath), 0755); err != nil {
		return errors.Wrap(errors.ErrSaveFile, err)
	}
	return os.WriteFile(filePath, data, 0600)
}

func downloadFile(ctx context.Context, api *tgbotapi.BotAPI, fileID string) string {
	logger := logging.FromContext(ctx).With().Str("fileID", fileID).Logger()

	logger.Info().Msg("downloading file")

	file, err := getFile(api, fileID)
	if err != nil {
		logger.Error().Err(err).Msg("error getting file")
		return ""
	}

	filePath := path.Join(config.Misc.TempDir, "telegram", "downloads", file.FilePath)

	req, err := createRequest(file.Link(api.Token))
	if err != nil {
		logger.Error().Err(err).Msg("error creating request")
		return ""
	}

	data, err := downloadFileData(api, req)
	if err != nil {
		logger.Error().Err(err).Msg("error getting file URL")
		return ""
	}

	if err := saveFile(filePath, data); err != nil {
		logger.Error().Err(err).Msg("error saving file")
		return ""
	}

	if path.Ext(filePath) == ".ogg" || path.Ext(filePath) == ".oga" {
		nfilePath, err := convertAudioToMp3(ctx, filePath)
		defer func(name string) {
			_ = os.Remove(name)
		}(filePath)
		if err != nil {
			logger.Error().Err(err).Msg("error converting file")
			return ""
		}
		return nfilePath
	}

	return filePath
}
