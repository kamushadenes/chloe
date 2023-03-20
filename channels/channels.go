package channels

import (
	"github.com/kamushadenes/chloe/messages"
	"github.com/kamushadenes/chloe/structs"
)

var (
	IncomingMessagesCh   = make(chan *messages.Message, 100)
	OutgoingMessagesCh   = make(chan *OutgoingMessage, 100)
	CompletionRequestsCh = make(chan *structs.CompletionRequest)
	TranscribeRequestsCh = make(chan *structs.TranscriptionRequest)
	GenerationRequestsCh = make(chan *structs.GenerationRequest)
	EditRequestsCh       = make(chan *structs.GenerationRequest)
	VariationRequestsCh  = make(chan *structs.VariationRequest)
	TTSRequestsCh        = make(chan *structs.TTSRequest)
)
