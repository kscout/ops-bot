package messages

// Message from a user to the bot
//
// It is assumed that if a message is turned into a Message struct the
// user who sent the message has permission to interact with the bot.
// It is up to the specific message source to verify this.
type Message struct {
	// Body of message
	Body string

	// Responder which will respond to the source message
	Responder Responder
}
