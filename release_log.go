package contracts

import "time"

// ReleaseLog represents a release log for a specific release
type ReleaseLog struct {
	ID         string          `json:"id,omitempty"`
	RepoSource string          `json:"repoSource"`
	RepoOwner  string          `json:"repoOwner"`
	RepoName   string          `json:"repoName"`
	ReleaseID  string          `json:"releaseID"`
	Steps      []*BuildLogStep `json:"steps"`
	InsertedAt time.Time       `json:"insertedAt"`
}

// GetAggregatedStatus returns the status aggregated across all stages
func (releaseLog *ReleaseLog) GetAggregatedStatus() LogStatus {
	return GetAggregatedStatus(releaseLog.Steps)
}

// HasUnknownStatus returns true if aggregated status is unknown
func (releaseLog *ReleaseLog) HasUnknownStatus() bool {
	return HasUnknownStatus(releaseLog.Steps)
}

// HasSucceededStatus returns true if aggregated status is succeeded
func (releaseLog *ReleaseLog) HasSucceededStatus() bool {
	return HasSucceededStatus(releaseLog.Steps)
}

// HasFailedStatus returns true if aggregated status is failed
func (releaseLog *ReleaseLog) HasFailedStatus() bool {
	return HasFailedStatus(releaseLog.Steps)
}

// HasCanceledStatus returns true if aggregated status is canceled
func (releaseLog *ReleaseLog) HasCanceledStatus() bool {
	return HasSucceededStatus(releaseLog.Steps)
}
