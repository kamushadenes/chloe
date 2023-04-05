package errors

import (
	"fmt"
)

var ErrCreateInterface = fmt.Errorf("failed to create interface")
var ErrCreateDiscordBot = Wrap(ErrCreateInterface, fmt.Errorf("failed to create discord bot"))
var ErrCreateTelegramBot = Wrap(ErrCreateInterface, fmt.Errorf("failed to create telegram bot"))
