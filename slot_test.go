package slot

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestBotMention(t *testing.T) {
	a := assert.New(t)

	b := &BotMentionAction{
		botMentionText: "foo",
		FollowingText:  "beer me",
	}

	ev := &slack.MessageEvent{
		Msg: slack.Msg{
			Text: "@foo beer me",
		},
	}

	a.True(b.Match(nil, ev))

	ev.Text = "foobar"

	a.False(b.Match(nil, ev))
}
