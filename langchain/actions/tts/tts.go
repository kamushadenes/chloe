package tts

import (
	"fmt"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/base"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *TTSAction) GetNotification() string {
	return fmt.Sprintf("🔉 Generating audio: **%s**", a.MustGetParam("text"))
}

func (a *TTSAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Audio)

	tts := base.NewTTSWithDefaultModel(config.TTS.Provider)

	res, err := tts.TTSWithContext(request.Context, common.TTSMessage{Text: a.MustGetParam("text")})
	if err != nil {
		return nil, err
	}

	obj.Header().Add("Content-Type", res.ContentType)

	_, err = obj.Write(res.Audio)

	return []*response_object_structs.ResponseObject{obj}, err
}
