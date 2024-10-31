package contracts

import (
	"fmt"
	"time"

	manifest "github.com/estafette/estafette-ci-manifest"
)

// Pipeline represents a pipeline with the latest build info, including version number, repo, branch, revision, labels and manifest
type Pipeline struct {
	ID                   string                      `json:"id"`
	RepoSource           string                      `json:"repoSource"`
	RepoOwner            string                      `json:"repoOwner"`
	RepoName             string                      `json:"repoName"`
	RepoBranch           string                      `json:"repoBranch"`
	RepoRevision         string                      `json:"repoRevision"`
	BuildVersion         string                      `json:"buildVersion,omitempty"`
	BuildStatus          Status                      `json:"buildStatus,omitempty"`
	Labels               []Label                     `json:"labels,omitempty"`
	ReleaseTargets       []ReleaseTarget             `json:"releaseTargets,omitempty"`
	Manifest             string                      `json:"manifest,omitempty"`
	ManifestWithDefaults string                      `json:"manifestWithDefaults,omitempty"`
	Commits              []GitCommit                 `json:"commits,omitempty"`
	Triggers             []manifest.EstafetteTrigger `json:"triggers,omitempty"`
	Events               []manifest.EstafetteEvent   `json:"triggerEvents,omitempty"`
	Archived             bool                        `json:"archived,omitempty"`
	InsertedAt           time.Time                   `json:"insertedAt"`
	StartedAt            *time.Time                  `json:"startedAt,omitempty"`
	UpdatedAt            time.Time                   `json:"updatedAt"`
	Duration             time.Duration               `json:"duration"`
	PendingDuration      *time.Duration              `json:"pendingDuration,omitempty"`
	LastUpdatedAt        time.Time                   `json:"lastUpdatedAt"`
	ManifestObject       *manifest.EstafetteManifest `json:"-"`
	RecentCommitters     []string                    `json:"recentCommitters,omitempty"`
	RecentReleasers      []string                    `json:"recentReleasers,omitempty"`
	ExtraInfo            *PipelineExtraInfo          `json:"extraInfo,omitempty"`
	Groups               []*Group                    `json:"groups,omitempty"`
	Organizations        []*Organization             `json:"organizations,omitempty"`
}

// GetFullRepoPath returns the full path of the pipeline repository with source, owner and name
func (pipeline *Pipeline) GetFullRepoPath() string {
	return fmt.Sprintf("%v/%v/%v", pipeline.RepoSource, pipeline.RepoOwner, pipeline.RepoName)
}

// PipelineExtraInfo contains extra information like aggregates over the last x builds
type PipelineExtraInfo struct {
	MedianPendingDuration time.Duration `json:"medianPendingDuration"`
	MedianDuration        time.Duration `json:"medianDuration"`
}
