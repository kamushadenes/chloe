package youtube

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	errors2 "github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/react/actions/transcribe"
	"github.com/kamushadenes/chloe/structs"
	"os"
	"os/exec"
	"path"
)

type YoutubeTranscribeAction struct {
	Name   string
	Params map[string]string
}

func (a *YoutubeTranscribeAction) GetNotification() string {
	return fmt.Sprintf("🎥️ Transcribing video: **%s** (this might take a while)", a.Params["url"])
}

func (a *YoutubeTranscribeAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	logger := logging.FromContext(request.Context).With().Str("action", a.GetName()).Str("url", a.Params["url"]).Logger()

	obj := structs.NewResponseObject(structs.Text)

	if _, err := exec.LookPath("youtube-dl"); err != nil {
		return nil, errors2.Wrap(errors2.ErrActionFailed, errors2.ErrCommandNotFound, err)
	}

	tmpDir, err := os.MkdirTemp(config.Misc.TempDir, "youtube")
	if err != nil {
		return nil, errors2.Wrap(errors2.ErrActionFailed, err)
	}

	logger.Info().Msg("downloading audio")

	var args []string
	if config.React.UseAria2 {
		if _, err := exec.LookPath("aria2c"); err == nil {
			args = append(args, "--external-downloader", "aria2c")
		} else {
			logger.Warn().Err(err).Msg("aria2c not found, falling back to default downloader")
		}
	}
	args = append(args,
		"-x", "--audio-format", "mp3",
		"-f", "worstaudio/bestaudio/worst",
		"-o", path.Join(tmpDir, "audio.mp3"),
		a.Params["url"])

	cmd := exec.Command("youtube-dl", args...)
	if err := cmd.Run(); err != nil {
		return nil, errors2.Wrap(errors2.ErrActionFailed, errors2.ErrCommandError, err)
	}

	na := transcribe.NewTranscribeAction()
	na.SetParam("path", path.Join(tmpDir, "audio.mp3"))
	na.SetMessage(request.Message)
	request.Message.NotifyAction(na.GetNotification())

	tobjs, err := na.Execute(request)
	if err != nil {
		return nil, errors2.Wrap(errors2.ErrActionFailed, err)
	}

	for k := range tobjs {
		if err := obj.WriteObject(tobjs[k]); err != nil {
			return nil, errors2.Wrap(errors2.ErrActionFailed, err)
		}
	}

	return []*structs.ResponseObject{obj}, nil
}
