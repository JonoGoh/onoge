package commands

import (
	"github.com/bwmarrin/discordgo"
	"onoge/commands/slash"
)

// HandleSlashCommand processes slash commands sent to the bot
func HandleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.ApplicationCommandData().Name {
	case "lonnie":
		slash.HandleLonnieCommand(s, i)
	}
}