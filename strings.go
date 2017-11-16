package slot

import "github.com/nlopes/slack"

// TextResponder always responds with a string
type TextResponder string

// Respond implements Responder
func (p TextResponder) Respond(r *slack.RTM, ev *slack.MessageEvent) error {
	r.SendMessage(&slack.OutgoingMessage{
		Channel: ev.Channel,
		Text:    string(p),
		Type:    slack.TYPE_MESSAGE,
	})
	return nil
}

// StringFuncResponder always responds with the result of calling the function
type StringFuncResponder func() string

// Respond implements Responder
func (s StringFuncResponder) Respond(r *slack.RTM, ev *slack.MessageEvent) error {
	r.SendMessage(&slack.OutgoingMessage{
		Channel: ev.Channel,
		Text:    s(),
		Type:    slack.TYPE_MESSAGE,
	})
	return nil
}
