package channels

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func RegisterIncomingMessage(msg *memory.Message) error {
	if msg.ErrorCh == nil {
		msg.ErrorCh = make(chan error)
	}
	IncomingMessagesCh <- msg

	if err := <-msg.ErrorCh; err != nil {
		return errors.Wrap(errors.ErrSavingMessage, err)
	}

	return nil
}

func RunCompletion(req *structs.CompletionRequest) error {
	if req.ErrorChannel == nil {
		req.ErrorChannel = make(chan error)
	}
	CompletionRequestsCh <- req

	if err := <-req.ErrorChannel; err != nil {
		return errors.Wrap(errors.ErrCompletionFailed, err)
	}

	return nil
}

func RunAction(req *structs.ActionRequest) error {
	if req.ErrorChannel == nil {
		req.ErrorChannel = make(chan error)
	}
	ActionRequestsCh <- req

	if err := <-req.ErrorChannel; err != nil {
		return errors.Wrap(errors.ErrActionFailed, fmt.Errorf("error running action %s", req.Action), err)
	}

	return nil
}
