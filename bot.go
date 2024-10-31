package contracts

import (
	"fmt"
	"time"

	manifest "github.com/estafette/estafette-ci-manifest"
)

// Bot represents a bot execution
type Bot struct {
	Name            string                    `json:"name"`
	ID              string                    `json:"id,omitempty"`
	RepoSource      string                    `json:"repoSource,omitempty"`
	RepoOwner       string                    `json:"repoOwner,omitempty"`
	RepoName        string                    `json:"repoName,omitempty"`
	BotStatus       Status                    `json:"botStatus,omitempty"`
	Events          []manifest.EstafetteEvent `json:"triggerEvents,omitempty"`
	InsertedAt      *time.Time                `json:"insertedAt,omitempty"`
	StartedAt       *time.Time                `json:"startedAt,omitempty"`
	UpdatedAt       *time.Time                `json:"updatedAt,omitempty"`
	Duration        *time.Duration            `json:"duration,omitempty"`
	PendingDuration *time.Duration            `json:"pendingDuration,omitempty"`
	ExtraInfo       *BotExtraInfo             `json:"extraInfo,omitempty"`
	Groups          []*Group                  `json:"groups,omitempty"`
	Organizations   []*Organization           `json:"organizations,omitempty"`
}

// BotExtraInfo contains extra information like aggregates over the last x releases
type BotExtraInfo struct {
	MedianPendingDuration time.Duration `json:"medianPendingDuration"`
	MedianDuration        time.Duration `json:"medianDuration"`
}

// GetFullRepoPath returns the full path of the bot repository with source, owner and name
func (bot *Bot) GetFullRepoPath() string {
	return fmt.Sprintf("%v/%v/%v", bot.RepoSource, bot.RepoOwner, bot.RepoName)
}
