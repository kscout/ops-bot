[![GoDoc](https://godoc.org/github.com/kscout/ops-bot?status.svg)](https://godoc.org/github.com/kscout/ops-bot)

# Ops Bot
Development helper bot.

# Table Of Contents
- [Overview](#overview)
- [Interact](#interact)
  - [Summon](#summon)
  - [Permissions](#permissions)
  - [Commands](#commands)
	- [Deploy Command](#deploy-command)
- [Configure](#configure)
  - [Configure GitHub App](#configure-github-app)
  - [Configure Application](#configure-application)
- [Develop](#develop)
  
# Overview
Chat bot which completes development operations.

See [interact section](#interact) for details on how to use the bot.

See the [develop section](#develop) to learn how to contribute.

# Interact
## Summon
Bot responds to GitHub pull request messages in the KScout organization.

Bot only takes action if summoned via an @-mention:

```
@kscout-ops-bot
```

Summon can be anywhere in comment.

## Permissions
Only users which are part of the KScout GitHub organization can interact with 
the bot.

## Commands
Accepts commands in a rigid syntax:

```
/<command> <options>
```

Comments can have an unlimited number of commands, only requirement is that all
commands start with a slash.

### Deploy Command
Deploy pull request to environment.

Usage:

```
/deploy env=<env>
```

Options:

- `<env>` (String) Environment to deploy pull request
  
# Configure
## Configure GitHub App
Create a new GitHub app:

1. Set the Homepage URL to `https://github.com/kscout/ops-bot`
2. Set the Webhook URL to `https://<deploy host>/github/webhook`
3. Set the Webhook Secret
4. Permissions:
   - Repository administration: Read-only
   - Repository contents: Read-only
   - Deployments: Read & write
   - Issues: Read & write
   - Pull requests: Read & write
5. Subscribe to events
   - Commit comment
   - Issue comment
   - Pull request
   - Pull request review
   - Pull request review comment
   - Push
6. Click create button
7. Set application logo

Install GitHub app:

1. Install on all repositories
2. Generate private key

## Configure Application
Configuration is passed via environment variables.

See the [config Godoc](https://godoc.org/github.com/kscout/ops-bot/config) for 
information about the available configuration options.

# Develop
Most of the developer documentation is located in the 
[godoc](https://godoc.org/github.com/kscout/ops-bot]. See the 
[handlers godoc](https://godoc.org/github.com/kscout/ops-bot/handlers] for 
HTTP API documentation.
