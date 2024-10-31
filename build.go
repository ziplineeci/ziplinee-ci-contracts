package contracts

import (
	"fmt"
	"time"

	manifest "github.com/estafette/estafette-ci-manifest"
)

// Build represents a specific build, including version number, repo, branch, revision, labels and manifest
type Build struct {
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
	InsertedAt           time.Time                   `json:"insertedAt"`
	StartedAt            *time.Time                  `json:"startedAt,omitempty"`
	UpdatedAt            time.Time                   `json:"updatedAt"`
	Duration             time.Duration               `json:"duration"`
	PendingDuration      *time.Duration              `json:"pendingDuration,omitempty"`
	ManifestObject       *manifest.EstafetteManifest `json:"-"`
	Groups               []*Group                    `json:"groups,omitempty"`
	Organizations        []*Organization             `json:"organizations,omitempty"`
}

// GetFullRepoPath returns the full path of the build repository with source, owner and name
func (build *Build) GetFullRepoPath() string {
	return fmt.Sprintf("%v/%v/%v", build.RepoSource, build.RepoOwner, build.RepoName)
}
