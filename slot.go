package slot

import (
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

type ConditionalAction interface {
	Match(*slack.RTM, *slack.MessageEvent) bool
	Execute(*slack.RTM, *slack.MessageEvent) error
}

type MatchFunc func(*slack.RTM, *slack.MessageEvent) bool
type ActionFunc func(*slack.RTM, *slack.MessageEvent) error

type regexAction struct {
	re     *regexp.Regexp
	action ActionFunc
}

func (r *regexAction) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool {
	return r.re.MatchString(ev.Text)
}

func (r *regexAction) Execute(rtm *slack.RTM, ev *slack.MessageEvent) error {
	return r.action(rtm, ev)
}

type BotMentionAction struct {
	botMentionText string
	FollowingText  string
	Action         ActionFunc
}

func (b *BotMentionAction) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool {
	if b.botMentionText == "" {
		if id, err := rtm.GetUserIdentity(); err == nil {
			b.botMentionText = "@" + id.User.Name
		}
	}

	return strings.Contains(ev.Text, b.botMentionText+" "+b.FollowingText)
}

func (b *BotMentionAction) Execute(rtm *slack.RTM, ev *slack.MessageEvent) error {
	return b.Action(rtm, ev)
}
