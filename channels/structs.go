package channels

import (
	"github.com/kamushadenes/chloe/users"
	"io"
)

type OutgoingMessage struct {
	User         *users.User
	Texts        []string
	Audios       []string
	Images       []string
	TextWriters  []io.WriteCloser
	AudioWriters []io.WriteCloser
	ImageWriters []io.WriteCloser
}
