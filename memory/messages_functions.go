package memory

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/sashabaranov/go-openai"
)

func LoadNonSummarizedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("summary IS NULL AND content IS NOT NULL").Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func LoadNonModeratedMessages(ctx context.Context) ([]*Message, error) {
	var messages []*Message

	if err := db.WithContext(ctx).Where("moderated = ? AND "+
		"content IS NOT NULL AND "+
		"content != ''", false).
		Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func MessagesFromOpenAIChatCompletionResponse(ctx context.Context, user *User, interf string, resp *openai.ChatCompletionResponse) []*Message {
	var messages []*Message

	for k := range resp.Choices {
		msg := NewMessage(uuid.Must(uuid.NewV4()).String(), interf)

		msg.Content = resp.Choices[k].Message.Content
		msg.Role = resp.Choices[k].Message.Role
		msg.User = user

		messages = append(messages, msg)
	}

	return messages
}
