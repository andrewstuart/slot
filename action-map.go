package slot

import (
	"strings"

	"log"

	"github.com/nlopes/slack"
)

// GetAction takes an event and returns either the empty string, or the first
// !action string in the message text.
func GetAction(ev *slack.MessageEvent) string {
	if ev.Type != slack.TYPE_MESSAGE {
		return ""
	}
	key := strings.Split(ev.Text, " ")[0]
	if len(key) < 1 || key[0] != '!' {
		return ""
	}

	return strings.TrimPrefix(key, "!")
}

// ActionMap holds action words and responders, calling the appropriate
// responder when an !action message is received.
type ActionMap map[string]Responder

// Match implements MatchResponder for efficiency
func (m ActionMap) Match(r *slack.RTM, ev *slack.MessageEvent) bool {
	action := GetAction(ev)
	if action == "" {
		return false
	}
	log.Println(action)

	// Return whether map contains the action
	_, ok := m[action]
	return ok
}

// Respond implements Responder
func (m ActionMap) Respond(r *slack.RTM, ev *slack.MessageEvent) error {
	return m[GetAction(ev)].Respond(r, ev)
}
