package contracts

import "time"

// BotLog represents a bot log for a specific bot execution
type BotLog struct {
	ID         string          `json:"id,omitempty"`
	RepoSource string          `json:"repoSource"`
	RepoOwner  string          `json:"repoOwner"`
	RepoName   string          `json:"repoName"`
	BotID      string          `json:"botID"`
	Steps      []*BuildLogStep `json:"steps"`
	InsertedAt time.Time       `json:"insertedAt"`
}

// GetAggregatedStatus returns the status aggregated across all stages
func (botLog *BotLog) GetAggregatedStatus() LogStatus {
	return GetAggregatedStatus(botLog.Steps)
}

// HasUnknownStatus returns true if aggregated status is unknown
func (botLog *BotLog) HasUnknownStatus() bool {
	return HasUnknownStatus(botLog.Steps)
}

// HasSucceededStatus returns true if aggregated status is succeeded
func (botLog *BotLog) HasSucceededStatus() bool {
	return HasSucceededStatus(botLog.Steps)
}

// HasFailedStatus returns true if aggregated status is failed
func (botLog *BotLog) HasFailedStatus() bool {
	return HasFailedStatus(botLog.Steps)
}

// HasCanceledStatus returns true if aggregated status is canceled
func (botLog *BotLog) HasCanceledStatus() bool {
	return HasSucceededStatus(botLog.Steps)
}
