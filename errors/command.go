package errors

import (
	"fmt"
)

var ErrCommandError = fmt.Errorf("command error")
var ErrFFmpegError = Wrap(ErrCommandError, fmt.Errorf("ffmpeg error"))
var ErrImageMagickError = Wrap(ErrCommandError, fmt.Errorf("imagemagick error"))
var ErrCommandNotFound = Wrap(ErrCommandError, fmt.Errorf("command not found"))
