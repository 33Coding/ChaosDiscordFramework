package server

import (
	"github.com/bwmarrin/discordgo"
  "github.com/jaredfolkins/badactor"
)

type ChaosDiscord struct {
	Disc *discordgo.Session
  BadActor *badactor.Studio
}
