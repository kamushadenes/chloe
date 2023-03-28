package channels

import (
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

var (
	IncomingMessagesCh   = make(chan *memory.Message, 100)
	CompletionRequestsCh = make(chan *structs.CompletionRequest, 10)
	TranscribeRequestsCh = make(chan *structs.TranscriptionRequest, 10)
	GenerationRequestsCh = make(chan *structs.GenerationRequest, 10)
	EditRequestsCh       = make(chan *structs.GenerationRequest, 10)
	VariationRequestsCh  = make(chan *structs.VariationRequest, 10)
	TTSRequestsCh        = make(chan *structs.TTSRequest, 10)
	ActionRequestsCh     = make(chan *structs.ActionRequest, 10)
)
