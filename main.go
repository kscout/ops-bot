package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/kscout/ops-bot/commands"
	"github.com/kscout/ops-bot/config"
	"github.com/kscout/ops-bot/handlers"
	"github.com/kscout/ops-bot/messages"

	"github.com/Noah-Huppert/golog"
)

func main() {
	// {{{1 Initial setup
	context, cancelFn := context.WithCancel(context.Background())
	logger := golog.NewStdLogger("ops-bot")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs
		cancelFn()
	}()

	// doneGroup is used to wait for all go routines started in main to finish before exiting
	// .Add should be called before a go routine is called, .Done should be called when the
	// go routine ends.
	var doneGroup sync.WaitGroup

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("failed to load configuration: %s", err.Error())
	}

	// {{{1 Setup message processing system
	var messageBus chan messages.Message
	var commandReqBus chan commands.CommandRequest

	cmdReceiver := commands.Receiver{
		MessageBus:    messageBus,
		CommandReqBus: commandReqBus,
	}

	doneGroup.Add(1)
	go func() {
		defer doneGroup.Done()

		logger.Debug("starting command receiver")

		cmdReceiver.Receive(ctx)

		logger.Debug("stopped command receiver")
	}()

	// {{{1 Setup HTTP server
	httpRouter := handlers.NewRouter()
	httpRouter.Add(handlers.HealthHandler{})
	httpRouter.Add(handlers.GHWebhookHandler{
		GH:              nil,
		GHWebhookSecret: cfg.GHWebhookSecret,
		MessageBus:      messageBus,
	})

	server := http.Server{
		Addr: cfg.HTTPAddr,
		Handler: handlers.PanicHandler{
			Logger: logger.GetChild("panic"),
			Hander: httpRouter,
		},
	}

	doneGroup.Add(1)
	go func() {
		defer doneGroup.Done()

		logger.Debugf("starting HTTP server to listen on \"%s\"", cfg.HTTPAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("failed to start HTTP server: %s", err.Error())
		}
	}()

	doneGroup.Add(1)
	go func() {
		defer doneGroup.Done()

		<-context.Done()

		logger.Debug("closing HTTP server")
		if err := server.Close(); err != nil {
			logger.Fatalf("failed to close HTTP server: %s", err.Error())
		}
	}()

	doneGroup.Wait()
}
