package slot

import (
	"testing"

	"github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

func TestActionString(t *testing.T) {
	asrt := assert.New(t)
	tab := [][2]string{
		{"!foo", "foo"},
		{"!bar", "bar"},
		{" !foo", ""},
		{"hey there!", ""},
		{"", ""},
	}

	am := ActionMap{}

	for _, entry := range tab {
		ev := &slack.MessageEvent{Msg: slack.Msg{Type: slack.TYPE_MESSAGE, Text: entry[0]}}
		asrt.Equal(entry[1], GetAction(ev))

		asrt.False(am.Match(nil, ev))
		am[entry[0]] = TextResponder(entry[1])
		asrt.Equal(entry[1] != "", am.Match(nil, ev))
	}

	asrt.Equal("", GetAction(&slack.MessageEvent{Msg: slack.Msg{Type: slack.TYPE_IM, Text: "!foo"}}))
}
