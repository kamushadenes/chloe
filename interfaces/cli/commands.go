package cli

import (
	"context"
)

type Globals struct {
	Context context.Context
}

var Flags struct {
	Version VersionFlag `name:"version" help:"Print version information and quit"`

	Complete CompleteCmd `cmd:"complete" short:"c" help:"Complete a prompt"`
	Generate GenerateCmd `cmd:"generate" short:"g" help:"Generate an prompt"`
	TTS      TTSCmd      `cmd:"tts" short:"t" help:"Generate an audio from a prompt"`

	Action ActionCmd `cmd:"" help:"Performs an action"`

	Forget ForgetCmd `cmd:"forget" short:"f" help:"Forget all users"`

	CountTokens CountTokensCmd `cmd:"count-tokens" help:"Count tokens"`

	CreateUser CreateUserCmd `cmd:"create-user" help:"Create a new user"`
	DeleteUser DeleteUserCmd `cmd:"delete-user" help:"Delete a user"`

	ListUsers  ListUsersCmd  `cmd:"list-users" help:"List users"`
	MergeUsers MergeUsersCmd `cmd:"merge-users" help:"Merge users"`

	AddExternalID    AddExternalIDCmd    `cmd:"add-external-id" help:"Add external ID to user"`
	DeleteExternalID DeleteExternalIDCmd `cmd:"delete-external-id" help:"Delete external ID from user"`

	ListMessages ListMessagesCmd `cmd:"list-messages" help:"List messages"`

	CreateAPIKey CreateAPIKeyCmd `cmd:"create-api-key" help:"Create an API key for use with the HTTP interface"`
}
