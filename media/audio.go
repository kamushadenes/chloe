package media

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"os/exec"
)

func ConvertAudioToMp3(ctx context.Context, filePath string) (string, error) {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return "", fmt.Errorf("unable to locate `ffmpeg`: %w", err)
	}

	logger := logging.FromContext(ctx).With().Str("filePath", filePath).Logger()

	logger.Info().Msg("converting audio to mp3")

	npath := filePath + ".mp3"

	cmd := exec.Command("ffmpeg", "-i", filePath, npath)
	b, err := cmd.CombinedOutput()

	fmt.Println(string(b))

	if err != nil {
		return npath, errors.Wrap(errors.ErrFFmpegError, err)
	}

	return npath, nil
}
