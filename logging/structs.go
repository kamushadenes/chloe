package logging

import "github.com/rs/zerolog"

type ConsoleWriter struct {
	zerolog.ConsoleWriter
}
