package main

import (
	"net/http"
	"context"
	"os"
	"os/signal"
	"sync"
	
	"github.com/Noah-Huppert/golog"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

// Config for ops bot
type Config struct {
	// HTTPAddr is the address to start the HTTP API server
	HTTPAddr string `default:":5000" split_words:"true" required:"true"`
}

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

	var doneGroup sync.WaitGroup

	var cfg Config
	if err := envconfig.Process("app", &cfg); err != nil {
		logger.Fatalf("failed to load configuration: %s", err.Error())
	}

	// {{{1 Setup HTTP server
	router := mux.NewRouter()
	
	server := http.Server{
		Addr: cfg.HTTPAddr,
		Handler: router,
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
