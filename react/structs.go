package react

import (
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type Action interface {
	GetName() string
	GetNotification() string
	SetParams(string)
	GetParams() string
	Execute(*structs.ActionRequest) error
	SetWriters([]io.WriteCloser)
	GetWriters() []io.WriteCloser
	SetMessage(*memory.Message)
	RunPreActions(*structs.ActionRequest) error
	RunPostActions(*structs.ActionRequest) error
}

type ChainOfThoughtResponse struct {
	Thought string `json:"thought"`
	Action  string `json:"action"`
	Params  string `json:"params"`
}

type ExtractorAPIResponse struct {
	Url           string        `json:"url"`
	Status        string        `json:"status"`
	StatusCode    int           `json:"status_code"`
	Domain        string        `json:"domain"`
	Title         string        `json:"title"`
	Author        []interface{} `json:"author"`
	DatePublished interface{}   `json:"date_published"`
	Images        []string      `json:"images"`
	Videos        []interface{} `json:"videos"`
	Text          string        `json:"text"`
	Html          string        `json:"html"`
}
