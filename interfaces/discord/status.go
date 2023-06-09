package discord

import "github.com/bwmarrin/discordgo"

var activies = []*discordgo.Activity{
	{
		Name: "with data",
		Type: discordgo.ActivityTypeGame,
	},
	{
		Name: "with fire",
		Type: discordgo.ActivityTypeGame,
	},
	{
		Name: "your steps",
		Type: discordgo.ActivityTypeWatching,
	},
	{
		Name: "your commands",
		Type: discordgo.ActivityTypeListening,
	},
	{
		Name: "my heart",
		Type: discordgo.ActivityTypeListening,
	},
	{
		Name: "to be a good AI",
		Type: discordgo.ActivityTypeCompeting,
	},
}

func updateStatus(session *discordgo.Session) error {
	return session.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: activies,
	})
}
