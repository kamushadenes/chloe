package media

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"os/exec"
)

func ConvertImageToPNG(filePath string) (string, error) {
	if _, err := exec.LookPath("convert"); err != nil {
		return "", fmt.Errorf("unable to locate `convert`: %w", err)
	}

	npath := filePath + ".png"

	cmd := exec.Command("convert",
		"-background", "none",
		"-gravity", "center",
		"-resize", "1024x1024>",
		"-extent", "1:1>",
		filePath, npath)

	return npath, errors.Wrap(errors.ErrImageMagickError, cmd.Run())
}
