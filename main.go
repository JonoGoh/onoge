package main //test

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	// "github.com/joho/godotenv" // Import the godotenv package
	"onoge/commands"     // Import the commands package
)

func main() {
	// Load environment variables from the .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Load the token from the environment variables
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("DISCORD_TOKEN environment variable is not set.")
	}

	// Create a new Discord session
	discordConnection, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
	}

	// Register handlers
	discordConnection.AddHandler(ready)
	discordConnection.AddHandler(commands.HandleSlashCommand)

	// Open a WebSocket connection to Discord and begin listening
	err = discordConnection.Open()
	if err != nil {
		log.Fatalf("error opening connection: %v", err)
	}
	defer discordConnection.Close()

	// Gracefully shutdown the bot when receiving an interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	log.Println("Bot is shutting down.")
}

// ready is triggered when the bot is fully started and ready to operate
func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	commands.RegisterCommands(s) // This now calls the separate command registration logic
}