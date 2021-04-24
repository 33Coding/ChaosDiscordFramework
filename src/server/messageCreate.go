package server

import (
	"CozyBot/src/bot"
	"CozyBot/src/handler"
	"errors"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Commands = []bot.Command{{
	Id:         "getrank",
	Name:       "get rank",
	Ids:        []string{"r", "rank", "level", "exp"},
	Repeatable: true,
	Handler:    handler.RankHandler,
}, {
	Id:         "whoami",
	Name:       "know who he is",
	Ids:        []string{"whoami", "?"},
	Repeatable: false,
	Handler:    handler.WhoamiHandler,
}}

func (cb *CozyBot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := bot.Message{
		Session: s,
		Message: m,
	}

	if len(m.Content) > 1 &&
		(m.Content[0:1] == "!" || m.Content[0:1] == "." || m.Content[0:1] == ":" || m.Content[0:1] == "-") {

		message.Delete()

		//spam check
		if cb.BadActor.IsJailed(message.Message.Message.Author.ID) {
			log.Println("spam block")

			err := message.Session.ChannelMessageDelete(message.Message.ChannelID, message.Message.ID)
			if err != nil {
				log.Println(err)
			}
			return
		}

		err := cb.BadActor.Infraction(message.Message.Message.Author.ID, "spamBlock") // add an Infraction on every command usage (even if invalid one)
		if err != nil {
			log.Println(err)
		}

		message.UpdateCommandReply()

    if err != nil{
      return
    }
		command := strings.Split(m.Content[1:], " ")

		command[0] = strings.ToLower(command[0])

	SearchCmd:
		for _, cmd := range Commands {
			for i := range cmd.Ids {
				if cmd.Ids[i] == command[0] {
					message.Command = &cmd
					break SearchCmd
				}
			}
		}

		if message.Command == nil {
			message.Error = errors.New("⚠ invalid command")
			message.Closed = true
			message.UpdateCommandReply()
		} else {
			message.UpdateCommandReply()
			message.HandleCommand()
		}

		//commandHandler(&message)
	} else {
		//messageHandler(&message)
		err := cb.LogXp(s, m)

		if err != nil {
			log.Println("Error logging exp:")
			log.Println(err)
		}
	}

}
