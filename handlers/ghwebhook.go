package handlers

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"encoding/json"

	"github.com/kscout/ops-bot/messages"
	"github.com/kscout/ops-bot/gh"
	
	"github.com/gorilla/mux"
	"github.com/google/go-github/github"
)

// GHWebhookHandler receives GitHub event notifications from GitHub.
//
// The webhook parses these events and translates them into messages.Message structs.
// These are then sent to a message bus and processed.
//
// Currently only issue + pull request comments are valid message sources. Edits to
// comments will be ignored.
type GHWebhookHandler struct {
	// GH API client
	GH *github.Client
	
	// GHWebhookSecret is key used to sign an HMAC of GitHub webhook requests
	// See config.Config#GHWebhookSecret for more details.
	GHWebhookSecret string

	// MessageBus receives messages.Message items and processes them
	MessageBus chan messages.Message
}

// GHWebhookResponse provides some very basic status information about what the webhook did
// This information is never used by the GitHub API. Instead it can be used by us to debug
// the webhook.
type GHWebhookResponse struct {
	// OK indicates the webhook completed successfully
	OK bool `json:"ok"`

	// HadMessages indicates if the event sent to the webhook had any messages
	HadMessages bool `json:"had_messages"`
}

// Register handler at POST /github/webhook
func (h GHWebhookHandler) Register(router *mux.Router) {
	router.Handle("/github/webhook", h).Methods("POST")
}

// ServeHTTP parses the GitHub event sent in the request and sends messages.Messages
// to the MessageBus.
//
// Responds with GHWebhookResponse.
func (h GHWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// {{{1 Verify request signature
	actualHubSig := r.Header().Get("X-Hub-Signature")
	
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(fmt.Errorf("failed to read request body: %s", err.Error()))
	}
	
	bodyHMAC := hmac.New(sha1.New, []byte(h.GHWebhookSecret))
	bodyHMAC.Write(bodyBytes)

	expectedHubSig := fmt.Sprintf("sha1=%s", hex.EncodeToString(bodyHMAC.Sum(nil)))

	if !hmac.Equal([]byte(expectedHubSig), []byte(actualHubSig)) {
		resp := JSONResponder{
			Status: http.StatusUnauthorized,
			Data: NewError(fmt.Errorf("request signature did not match expected value")),
		}
		resp.Respond(w)
		return
	}

	// {{{1 Spawn messages for events
	hadMessages := false
	
	switch r.Header().Get("X-Github-Event") {
	case "issue_comment":
		// {{{2 Parse
		var commentEvent github.IssueCommentEvent
		if err := json.Unmarshal(bodyBytes, &commentEvent); err != nil {
			panic(fmt.Errorf("failed to unmarshal issue comment event as JSON: %s",
				err.Error()))
		}

		// Ignore comment edits and deletes
		if commentEvent.GetAction() != "created" {
			break
		}

		// {{{2 Send message to bus
		h.MessageBus <- messages.Message{
			Body: commentEvent.GetComment().GetBody(),
			Responder: gh.IssueResponder{
				GH: h.GH,
				RepoOwner: commentEvent.Issue.Repo.Owner.GetLogin(),
				RepoName: commentEvent.Issue.Repo.GetName(),
				Number: commentEvent.Issue.GetNumber(),
			},
		}

		hadMessages = true
		
		break
	}

	// {{{1 Response
	resp := JSONResponder{
		Status: http.StatusOK,
		Data: GHWebhookResponse{
			OK: true,
			HadMessages: hadMessages,
		},
	}
	resp.Respond(w)
}
