package messages

import (
	"context"
)

// Responder sends a response message
type Responder interface {
	// Respond with Message
	Respond(ctx context.Context, msg Message) error
}
