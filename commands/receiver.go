package commands

import (
	"context"
	"strings"

	"github.com/kscout/ops-bot/messages"
)

// Receiver parses commands from Messages and outputs command start requests to
// the CommandRunner
//
// The Receive method should be started for the Receiver to function.
type Receiver struct {
	// MessageBus receives Messages from all message sources
	MessageBus chan messages.Message

	// CommandReqBus is sent CommandRequests by the Receiver
	CommandReqBus chan CommandRequest
}

// Receive loop and blocks for a message on MessageBus then parses it and sends a start command request
//
// Messages may have an unlimited number of commands in their body. Commands should start with a forward
// slash (/). The command name should come after the slash. Spaces are allowed between the slash and the
// command name. Options should follow the command name. Options are seperated from the command name and
// each other by spaces. An individual option should be in the format "<key>=<value>".
func (r Receiver) Receive(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-r.MessageBus:
			// cmdsTxt is an array of commands. Each item is string of command token.
			cmdsTxt := strings.Split(msg.Body, "/")

			for _, txt := range cmdsTxt {
				toks := strings.Split(txt, " ")

				if len(toks) == 0 {
					continue
				}

				cmdReq := CommandRequest{
					Name:    toks[0],
					Options: map[string]string{},
				}

				for _, opt := range toks[1:] {
					parts := strings.Split(opt, "=")

					key := parts[0]
					value := strings.Join(parts[1:], "=")

					cmdReq.Options[key] = value
				}

				r.CommandReqBus <- cmdReq
			}
		}
	}
}
