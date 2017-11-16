package slot

import (
	"fmt"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

// A Bot handles a client
type Bot struct {
	Responders []Responder

	botID string
}

// Handle manages an RTM based on the configured Handlers
func (b *Bot) Handle(cli *slack.Client) error {
	if b.Responders == nil {
		return fmt.Errorf("bot had no configured Responders")
	}

	rtm := cli.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			b.botID = ev.Info.User.ID
			log.Debug("Connected!")
		case *slack.MessageEvent:
			if ev.BotID == b.botID || ev.User == b.botID {
				log.Debug("Not responding to our own message")
			}
			log.Debug("%s: %q", ev.User, ev.Text)

			for i := range b.Responders {
				err := MaybeRespond(rtm, ev, b.Responders[i])
				if err != nil {
					log.Error("bot handler error", err)
				}
			}
		case *slack.OutgoingErrorEvent:
			log.Error("Outgoing event error encountered", ev)
		case *slack.InvalidAuthEvent:
			return fmt.Errorf("Invalid Auth")
		case *slack.RTMError:
			return ev
		}
	}
}
