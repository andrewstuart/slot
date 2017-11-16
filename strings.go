package slot

import "github.com/nlopes/slack"

// PlainTextResponder always responds with a string
type PlainTextResponder string

// Respond implements Responder
func (p PlainTextResponder) Respond(r *slack.RTM, ev *slack.MessageEvent) error {
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
