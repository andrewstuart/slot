package slot

import (
	"regexp"

	"github.com/nlopes/slack"
)

// RegexResponder matches a regex against an incoming string and executes a
// response if a match occurred
type RegexResponder struct {
	Regexp    *regexp.Regexp
	Responder Responder
}

// Match implements MatchResponder
func (r *RegexResponder) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool {
	return r.Regexp.MatchString(ev.Text)
}

// Respond implements Responder
func (r *RegexResponder) Respond(rtm *slack.RTM, ev *slack.MessageEvent) error {
	return r.Responder.Respond(rtm, ev)
}
