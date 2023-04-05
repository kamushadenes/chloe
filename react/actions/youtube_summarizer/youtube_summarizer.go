package youtube_summarizer

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	errors2 "github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react/actions/transcribe"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	utils2 "github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
	"os/exec"
	"path"
	"strings"
)

type YoutubeSummarizerAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewYoutubeSummarizerAction() structs.Action {
	return &YoutubeSummarizerAction{
		Name: "youtube_summarizer",
	}
}

func (a *YoutubeSummarizerAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *YoutubeSummarizerAction) GetWriters() []io.WriteCloser {
	return a.Writers
}
func (a *YoutubeSummarizerAction) GetName() string {
	return a.Name
}

func (a *YoutubeSummarizerAction) GetNotification() string {
	return fmt.Sprintf("🎥️ Summarizing video: **%s** (this might take a while)", a.Params)
}

func (a *YoutubeSummarizerAction) SetParams(params string) {
	a.Params = params
}

func (a *YoutubeSummarizerAction) GetParams() string {
	return a.Params
}

func (a *YoutubeSummarizerAction) SetMessage(message *memory.Message) {}

func (a *YoutubeSummarizerAction) Execute(request *structs.ActionRequest) error {
	if _, err := exec.LookPath("youtube-dl"); err != nil {
		return errors2.Wrap(errors2.ErrActionFailed, errors2.ErrCommandNotFound, err)
	}

	logger := logging.GetLogger().With().Str("action", a.GetName()).Str("url", a.Params).Logger()

	tmpDir, err := os.MkdirTemp(config.Misc.TempDir, "youtube")
	if err != nil {
		return errors2.Wrap(errors2.ErrActionFailed, err)
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
		a.Params)

	cmd := exec.Command("youtube-dl", args...)
	if err := cmd.Run(); err != nil {
		return errors2.Wrap(errors2.ErrActionFailed, errors2.ErrCommandError, err)
	}

	b := &utils2.BytesWriter{}

	na := transcribe.NewTranscribeAction()
	na.SetParams(path.Join(tmpDir, "audio.mp3"))
	na.SetMessage(request.Message)
	na.SetWriters([]io.WriteCloser{b})
	request.Message.NotifyAction(na.GetNotification())
	if err := na.Execute(request); err != nil {
		return errors2.Wrap(errors2.ErrActionFailed, err)
	}

	resp, err := reactOpenAI.SimpleCompletionRequest(request.Context, "video_summarizer", string(b.Bytes))
	if err != nil {
		return errors2.Wrap(errors2.ErrActionFailed, err)
	}

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	for _, w := range a.Writers {
		if _, err := w.Write([]byte(content)); err != nil {
			return errors2.Wrap(errors2.ErrActionFailed, err)
		}

	}

	return nil
}

func (a *YoutubeSummarizerAction) RunPreActions(request *structs.ActionRequest) error {
	return errors2.ErrNotImplemented
}

func (a *YoutubeSummarizerAction) RunPostActions(request *structs.ActionRequest) error {
	return errors2.ErrNotImplemented
}
