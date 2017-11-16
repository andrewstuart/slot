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

	for _, entry := range tab {
		asrt.Equal(entry[1], GetAction(&slack.MessageEvent{Msg: slack.Msg{Type: slack.TYPE_MESSAGE, Text: entry[0]}}))
	}

	asrt.Equal("", GetAction(&slack.MessageEvent{Msg: slack.Msg{Type: slack.TYPE_IM, Text: "!foo"}}))
}
