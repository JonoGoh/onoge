package slash

import (
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

// Define your list of quotes
var quotes = []string{
	"Get in the bin.",
	"Eat my arse.",
	"Sorry when? When did I ask?",
	"Shut up you sket.",
	"That's absolutely diabolical.",
	"Why the fuck is Bradley here?",
	"Oh my diddly days.",
	"Suck your mum.",
	"This is why we can't have nice things.",
}

// handleLonnieCommand is the logic for the /lonnie command
func HandleLonnieCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	quote := quotes[rand.Intn(len(quotes))]

	// Respond with a random quote
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: quote,
		},
	})
}
