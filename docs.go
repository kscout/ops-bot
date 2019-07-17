/*
An API server which provides ops-bot functionality.

Due to the nature of running jobs only 1 ops-bot process should run at a time.
If more than one job ran at a time we would have to write distributed job management
logic. This is not required right now.

Architecture:

Messages received from any source -> commands.Receiver -> commands.CommandRunner

Message sources send received messages as messages.Message structs to commands.Receiver.

Currently GitHub is the only message source. The handlers package includes a webhook which receives
GitHub events and sends them to the commands.Receiver.

The commands.Receiver parses commands.CommandRequests out of incoming messages and sends these
to commands.CommandRunner.

commands.CommandRunner receives commands.CommandRequests and starts a go routines for each command request.
*/
package main
