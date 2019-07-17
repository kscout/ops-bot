package commands

// CommandRequest is a request to start a command
type CommandRequest struct {
	// Name of command
	Name string

	// Options passed to command
	// A key with an empty value is a flag style option set to true.
	Options map[string]string
}
