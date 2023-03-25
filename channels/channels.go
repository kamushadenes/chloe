package channels

import (
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

var (
	IncomingMessagesCh   = make(chan *memory.Message, 100)
	CompletionRequestsCh = make(chan *structs.CompletionRequest)
	TranscribeRequestsCh = make(chan *structs.TranscriptionRequest)
	GenerationRequestsCh = make(chan *structs.GenerationRequest)
	EditRequestsCh       = make(chan *structs.GenerationRequest)
	VariationRequestsCh  = make(chan *structs.VariationRequest)
	TTSRequestsCh        = make(chan *structs.TTSRequest)
)
