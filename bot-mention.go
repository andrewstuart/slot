package slot

import (
	"strings"

	"github.com/nlopes/slack"
)

// BotMentionAction executes a responder if the bot's name is @mentioned
type BotMentionAction struct {
	botMentionText string
	FollowingText  string
	Responder      Responder
}

// Match implements MatchResponder
func (b *BotMentionAction) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool {
	if b.botMentionText == "" {
		if id, err := rtm.GetUserIdentity(); err == nil {
			b.botMentionText = "@" + id.User.Name
		}
	}

	return strings.Contains(ev.Text, b.botMentionText+" "+b.FollowingText)
}

// Respond implements Responder
func (b *BotMentionAction) Respond(rtm *slack.RTM, ev *slack.MessageEvent) error {
	return b.Responder.Respond(rtm, ev)
}
