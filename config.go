package main

// Config for ops bot
type Config struct {
	// HTTPAddr is the address to start the HTTP API server
	HTTPAddr string `default:":5000" split_words:"true" required:"true"`
}
