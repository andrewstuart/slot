package slot

import (
	"fmt"

	"github.com/cloudflare/cfssl/log"
	"github.com/nlopes/slack"
)

// A Bot handles a client
type Bot struct {
	Handlers []Responder

	botID string
}

func (b *Bot) Handle(cli *slack.Client) error {
	rtm := cli.NewRTM()
	go rtm.ManageConnection()

	for {
		msg := <-rtm.IncomingEvents
		fmt.Printf("ev = %#v\n", msg.Data)
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			b.botID = ev.Info.User.ID

			log.Info("Connected!")
		case *slack.MessageEvent:
			if ev.BotID == b.botID || ev.User == b.botID {
				log.Debug("Not responding to our own message")
			}
			log.Debug("%s: %q", ev.User, ev.Text)

			for i := range b.Handlers {
				err := MaybeRespond(rtm, ev, b.Handlers[i])
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
