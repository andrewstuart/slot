package slot

import (
	"regexp"

	"github.com/nlopes/slack"
)

type RegexAction struct {
	re     *regexp.Regexp
	action Responder
}

func (r *RegexAction) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool {
	return r.re.MatchString(ev.Text)
}

func (r *RegexAction) Respond(rtm *slack.RTM, ev *slack.MessageEvent) error {
	return r.action.Respond(rtm, ev)
}
