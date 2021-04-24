package server

import (
	"fmt"
  "github.com/gofiber/fiber/v2"


	"github.com/bwmarrin/discordgo"
)

func (ccc *ChaosDiscord) Run() {

	// Register the messageCreate func as a callback for MessageCreate events.
	ccc.Disc.AddHandler(ccc.messageCreate)

	// In this example, we only care about receiving message events.
	ccc.Disc.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err := ccc.Disc.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
  defer ccc.Disc.Close()
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("trust is a weakness")
    })

    app.Listen(":8000")

	// Cleanly close down the Discord session.
	
}
