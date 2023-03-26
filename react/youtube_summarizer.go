package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/sashabaranov/go-openai"
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

func NewYoutubeSummarizerAction() Action {
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
	logger := logging.GetLogger().With().Str("action", a.GetName()).Str("url", a.Params).Logger()

	tmpDir, err := os.MkdirTemp(config.Misc.TempDir, "youtube")
	if err != nil {
		return err
	}

	logger.Info().Msg("downloading audio")

	cmd := exec.Command("youtube-dl",
		"--external-downloader", "aria2c",
		"-x", "--audio-format", "mp3",
		"-f", "worstaudio/bestaudio/worst",
		"-o", path.Join(tmpDir, "audio.mp3"),
		a.Params)
	if err := cmd.Run(); err != nil {
		return err
	}

	b := &BytesWriter{}

	na := NewTranscribeAction()
	na.SetParams(path.Join(tmpDir, "audio.mp3"))
	na.SetMessage(request.Message)
	na.SetWriters([]io.WriteCloser{b})
	request.Message.NotifyAction(na.GetNotification())
	if err := na.Execute(request); err != nil {
		return err
	}

	prompt, err := resources.GetPrompt("video_summarizer", &resources.PromptArgs{
		Args: map[string]interface{}{},
		Mode: "video_summarizer",
	})

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.ChainOfThought,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: string(b.Bytes),
			},
		},
	}

	var resp openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(request.Context, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(request.Context, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting prompt improvement")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return err
	}

	resp = respi.(openai.ChatCompletionResponse)

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	for _, w := range a.Writers {
		if _, err := w.Write([]byte(content)); err != nil {
			return err
		}

	}

	return nil
}

func (a *YoutubeSummarizerAction) RunPreActions(request *structs.ActionRequest) error {
	return defaultPreActions(a, request)
}

func (a *YoutubeSummarizerAction) RunPostActions(request *structs.ActionRequest) error {
	return defaultPostActions(a, request)
}
