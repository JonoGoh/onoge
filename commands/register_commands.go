package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// RegisterCommands registers all the slash commands for the bot
func RegisterCommands(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "lonnie",
			Description: "Get a random inspirational quote!",
		},
		// You can add more commands here as needed
	}

	for _, cmd := range commands {
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", cmd)
		if err != nil {
			log.Fatalf("Cannot create slash command %s: %v", cmd.Name, err)
		}
	}
}
