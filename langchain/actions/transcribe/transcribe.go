package transcribe

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/audio_models/asr"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *TranscribeAction) GetNotification() string {
	return "✍️ Transcribing..."
}

func (a *TranscribeAction) SetMessage(message *memory.Message) {}

func (a *TranscribeAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	stt := asr.NewASRWithDefaultModel(config.STT.Provider)

	res, err := stt.Transcribe(a.MustGetParam("path"))
	if err != nil {
		return nil, err
	}

	if _, err := obj.Write([]byte(res.Text)); err != nil {
		return nil, err
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}

func (a *TranscribeAction) RunPostActions(request *action_structs.ActionRequest) error {
	return nil
}
