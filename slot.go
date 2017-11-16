package slot

import (
	"github.com/nlopes/slack"
)

// A MatchResponder conditionally acts on a message
type MatchResponder interface {
	Responder
	Match(*slack.RTM, *slack.MessageEvent) bool
}

// A Responder handles an event
type Responder interface {
	Respond(*slack.RTM, *slack.MessageEvent) error
}

// An ResponderFunc is a function that can respond to a slack event
type ResponderFunc func(*slack.RTM, *slack.MessageEvent) error

// Respond implements Responder.
func (f ResponderFunc) Respond(r *slack.RTM, ev *slack.MessageEvent) error {
	return f(r, ev)
}

// MaybeRespond checks if a Responder is a MatchResponder and conditionally
// exits if there is no match.
func MaybeRespond(r *slack.RTM, ev *slack.MessageEvent, res Responder) error {
	if mr, ok := res.(MatchResponder); ok && !mr.Match(r, ev) {
		return nil
	}
	return res.Respond(r, ev)
}
