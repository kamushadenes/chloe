package media

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"os/exec"
	"strconv"
	"time"
)

func GetMediaDuration(path string) (time.Duration, error) {
	if _, err := exec.LookPath("ffprobe"); err != nil {
		return 0, errors.Wrap(errors.ErrCommandNotFound, fmt.Errorf("unable to locate `ffprobe`"), err)
	}

	cmd := exec.Command("ffprobe",
		"-i", path,
		"-show_entries",
		"format=duration",
		"-v", "quiet",
		"-of", "csv=p=0",
	)
	out, err := cmd.Output()
	if err != nil {
		return 0, errors.Wrap(errors.ErrFFmpegError, err)
	}

	f, err := strconv.ParseFloat(string(out), 64)
	if err != nil {
		return 0, errors.Wrap(errors.ErrCommandError, err)
	}

	dur := time.Duration(f * float64(time.Second))

	return dur, nil
}
