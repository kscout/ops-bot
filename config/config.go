package config

import (
	"fmt"
	
	"github.com/kelseyhightower/envconfig"
)

// Config for ops bot.
//
// Values are passed via environment variables.
//
// Environment variables will be prefixed by APP and be in capital
// underscore case (ex., APP_CAPITAL_UNDERSCORE_CASE).
//
// For example the HTTPAddr field is set by the APP_HTTP_ADDR environment variable.
type Config struct {
	// HTTPAddr is the address to start the HTTP API server
	HTTPAddr string `default:":5000" split_words:"true" required:"true"`

	// GHWebhookSecret is a secret key used to sign an HMAC of GitHub webhook requests
	// GitHub includes this HMAC in requests, we will recompute this HMAC using this key
	// and ensure that the included HMAC and our HMAC match.
	GHWebhookSecret string `split_words:"true" required:"true"`
}

// NewConfig loads Config values from environment variables
func NewConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("app", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %s", err.Error())
	}

	return &cfg, nil
}
