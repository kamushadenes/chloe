package youtube

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/kamushadenes/chloe/config"
	errors2 "github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/transcribe"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *YoutubeTranscribeAction) GetNotification() string {
	return fmt.Sprintf("🎥️ Transcribing video: **%s** (this might take a while)", a.MustGetParam("url"))
}

func (a *YoutubeTranscribeAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	logger := logging.FromContext(request.Context).With().Str("action", a.GetName()).Str("url", a.MustGetParam("url")).Logger()

	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

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
		a.MustGetParam("url"))

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

	return []*response_object_structs.ResponseObject{obj}, nil
}
