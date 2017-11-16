[![GoDoc](https://godoc.org/astuart.co/slot?status.svg)](https://godoc.org/astuart.co/slot)

# slot
--
    import "astuart.co/slot"

package slot gives some helpful abstractions over the nlopes/slack RTM
integrations. Most use cases are intended to be made easier. The common bot
response abstraction is the Responder. Many implementations will be created to
assist in most of the common bot use cases.

## Usage

#### func  GetAction

```go
func GetAction(ev *slack.MessageEvent) string
```
GetAction takes an event and returns either the empty string, or the first
!action string in the message text.

#### func  MaybeRespond

```go
func MaybeRespond(r *slack.RTM, ev *slack.MessageEvent, res Responder) error
```
MaybeRespond checks if a Responder is a MatchResponder and conditionally exits
if there is no match.

#### type ActionMap

```go
type ActionMap map[string]Responder
```

ActionMap holds action words and responders, calling the appropriate responder
when an !action message is received.

#### func (ActionMap) Match

```go
func (m ActionMap) Match(r *slack.RTM, ev *slack.MessageEvent) bool
```
Match implements MatchResponder for efficiency

#### func (ActionMap) Respond

```go
func (m ActionMap) Respond(r *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder

#### type Bot

```go
type Bot struct {
	Responders []Responder
}
```

A Bot handles a client

#### func (*Bot) Handle

```go
func (b *Bot) Handle(cli *slack.Client) error
```
Handle manages an RTM based on the configured Handlers

#### type BotMentionAction

```go
type BotMentionAction struct {
	FollowingText string
	Responder     Responder
}
```

BotMentionAction executes a responder if the bot's name is @mentioned

#### func (*BotMentionAction) Match

```go
func (b *BotMentionAction) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool
```
Match implements MatchResponder

#### func (*BotMentionAction) Respond

```go
func (b *BotMentionAction) Respond(rtm *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder

#### type MatchResponder

```go
type MatchResponder interface {
	Responder
	Match(*slack.RTM, *slack.MessageEvent) bool
}
```

A MatchResponder conditionally acts on a message

#### type PlainTextResponder

```go
type PlainTextResponder string
```

PlainTextResponder always responds with a string

#### func (PlainTextResponder) Respond

```go
func (p PlainTextResponder) Respond(r *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder

#### type RegexResponder

```go
type RegexResponder struct {
	Regexp    *regexp.Regexp
	Responder Responder
}
```

RegexResponder matches a regex against an incoming string and executes a
response if a match occurred

#### func (*RegexResponder) Match

```go
func (r *RegexResponder) Match(rtm *slack.RTM, ev *slack.MessageEvent) bool
```
Match implements MatchResponder

#### func (*RegexResponder) Respond

```go
func (r *RegexResponder) Respond(rtm *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder

#### type Responder

```go
type Responder interface {
	Respond(*slack.RTM, *slack.MessageEvent) error
}
```

A Responder handles an event

#### type ResponderFunc

```go
type ResponderFunc func(*slack.RTM, *slack.MessageEvent) error
```

An ResponderFunc is a function that can respond to a slack event

#### func (ResponderFunc) Respond

```go
func (f ResponderFunc) Respond(r *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder.

#### type StringFuncResponder

```go
type StringFuncResponder func() string
```

StringFuncResponder always responds with the result of calling the function

#### func (StringFuncResponder) Respond

```go
func (s StringFuncResponder) Respond(r *slack.RTM, ev *slack.MessageEvent) error
```
Respond implements Responder
