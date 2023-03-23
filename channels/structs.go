package channels

import (
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type OutgoingMessage struct {
	Interface    string
	User         *memory.User
	Texts        []string
	Audios       []string
	Images       []string
	TextWriters  []io.WriteCloser
	AudioWriters []io.WriteCloser
	ImageWriters []io.WriteCloser
}
