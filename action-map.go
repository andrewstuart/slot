package slot

import (
	"strings"

	"github.com/nlopes/slack"
)

func GetAction(ev *slack.MessageEvent) string {
	key := strings.Split(ev.Text, " ")[0]
	if key[0] != '!' {
		return ""
	}

	return strings.TrimPrefix(key, "!")
}

type ActionMap map[string]ActionFunc

func (m ActionMap) Match(r *slack.RTM, ev *slack.MessageEvent) bool {
	action := GetAction(ev)
	if action == "" {
		return false
	}

	// Return whether map contains the action
	_, ok := m[action]
	return ok
}

func (m ActionMap) Execute(r *slack.RTM, ev *slack.MessageEvent) error {
	return m[GetAction(ev)](r, ev)
}
