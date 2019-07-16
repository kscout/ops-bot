package gh

import (
	"context"
	"fmt"

	"github.com/kscout/ops-bot/messages"
	
	"github.com/google/go-github/github"
)

// IssueResponder responds to an issue (or pull request) via a comment
type IssueResponder struct {
	// GH client
	GH *github.Client

	// RepoOwner is the name of the owner of the repository to which the issue belongs
	RepoOwner string

	// RepoName is the name of the repository to which the issue belongs
	RepoName string

	// Number of the issue or pull request, a user facing ID provided by GitHub
	Number int
}

// Respond to GitHub issue via a comment
func (r IssueResponder) Respond(ctx context.Context, msg messages.Message) error {
	_, _, err := r.GH.Issues.CreateComment(ctx, r.RepoOwner, r.RepoName, r.Number,
		&github.IssueComment{
			Body: &msg.Body,
		})
	if err != nil {
		return fmt.Errorf("failed to create GitHub comment: %s", err.Error())
	}

	return nil
}
