package gpt4all

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	model_openai "github.com/kamushadenes/chloe/langchain/chat_models/openai"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/sashabaranov/go-openai"
)

func NewChatGPT4All(model *common.ChatModel, user *memory.User) common.Chat {
	cnf := openai.DefaultConfig("-")
	cnf.BaseURL = config.GPT4All.Endpoint
	client := openai.NewClientWithConfig(cnf)

	return &model_openai.ChatOpenAI{Client: client, Model: model, User: user}
}
