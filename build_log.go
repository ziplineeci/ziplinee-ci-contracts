package contracts

import (
	"time"
)

// BuildLog represents a build log for a specific revision
type BuildLog struct {
	ID           string          `json:"id,omitempty"`
	RepoSource   string          `json:"repoSource"`
	RepoOwner    string          `json:"repoOwner"`
	RepoName     string          `json:"repoName"`
	RepoBranch   string          `json:"repoBranch"`
	RepoRevision string          `json:"repoRevision"`
	BuildID      string          `json:"buildID"`
	Steps        []*BuildLogStep `json:"steps"`
	InsertedAt   time.Time       `json:"insertedAt"`
}

// BuildLogStep represents the logs for a single step of a pipeline
type BuildLogStep struct {
	Step         string                   `json:"step"`
	Depth        int                      `json:"depth,omitempty"`
	Image        *BuildLogStepDockerImage `json:"image"`
	RunIndex     int                      `json:"runIndex,omitempty"`
	Duration     time.Duration            `json:"duration"`
	LogLines     []BuildLogLine           `json:"logLines"`
	ExitCode     int64                    `json:"exitCode"`
	Status       LogStatus                `json:"status"`
	AutoInjected bool                     `json:"autoInjected,omitempty"`
	NestedSteps  []*BuildLogStep          `json:"nestedSteps,omitempty"`
	Services     []*BuildLogStep          `json:"services,omitempty"`
}

// BuildLogStepDockerImage represents info about the docker image used for a step
type BuildLogStepDockerImage struct {
	Name                   string        `json:"name"`
	Tag                    string        `json:"tag"`
	IsPulled               bool          `json:"isPulled"`
	ImageSize              int64         `json:"imageSize"`
	PullDuration           time.Duration `json:"pullDuration"`
	Error                  string        `json:"error,omitempty"`
	IsTrusted              bool          `json:"isTrusted,omitempty"`
	HasInjectedCredentials bool          `json:"hasInjectedCredentials,omitempty"`
}

// BuildLogLine has low level log information
type BuildLogLine struct {
	LineNumber int       `json:"line,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
	StreamType string    `json:"streamType"`
	Text       string    `json:"text"`
}

// TailLogLine returns a log line for streaming logs to gui during a build
type TailLogLine struct {
	Step         string                   `json:"step"`
	ParentStage  string                   `json:"parentStage,omitempty"`
	Type         LogType                  `json:"type"`
	Depth        int                      `json:"depth,omitempty"`
	RunIndex     int                      `json:"runIndex,omitempty"`
	LogLine      *BuildLogLine            `json:"logLine,omitempty"`
	Image        *BuildLogStepDockerImage `json:"image,omitempty"`
	Duration     *time.Duration           `json:"duration,omitempty"`
	ExitCode     *int64                   `json:"exitCode,omitempty"`
	Status       *LogStatus               `json:"status,omitempty"`
	AutoInjected *bool                    `json:"autoInjected,omitempty"`
}

// GetAggregatedStatus returns the status aggregated across all stages
func (buildLog *BuildLog) GetAggregatedStatus() LogStatus {
	return GetAggregatedStatus(buildLog.Steps)
}

// GetAggregatedStatus returns the status aggregated across all stages
func GetAggregatedStatus(steps []*BuildLogStep) LogStatus {

	// aggregate per stage in order to take retries into account
	statusPerStage := map[string]LogStatus{}
	for _, s := range steps {
		if s.Status == LogStatusCanceled {
			return LogStatusCanceled
		}

		// last status for a stage is leading
		statusPerStage[s.Step] = s.Status
	}

	// if any stage ended in failure, the aggregated status is failed as well
	aggregatedStatus := LogStatusUnknown
	for _, status := range statusPerStage {
		if status == LogStatusSucceeded && aggregatedStatus == LogStatusUnknown {
			aggregatedStatus = status
		}
		if status == LogStatusFailed {
			aggregatedStatus = LogStatusFailed
		}
	}

	return aggregatedStatus
}

// HasUnknownStatus returns true if aggregated status is unknown
func (buildLog *BuildLog) HasUnknownStatus() bool {
	return HasUnknownStatus(buildLog.Steps)
}

// HasUnknownStatus returns true if aggregated status is unknown
func HasUnknownStatus(steps []*BuildLogStep) bool {
	status := GetAggregatedStatus(steps)

	return status == LogStatusUnknown
}

// HasSucceededStatus returns true if aggregated status is succeeded
func (buildLog *BuildLog) HasSucceededStatus() bool {
	return HasSucceededStatus(buildLog.Steps)
}

// HasSucceededStatus returns true if aggregated status is succeeded
func HasSucceededStatus(steps []*BuildLogStep) bool {
	status := GetAggregatedStatus(steps)

	return status == LogStatusSucceeded
}

// HasFailedStatus returns true if aggregated status is failed
func (buildLog *BuildLog) HasFailedStatus() bool {
	return HasFailedStatus(buildLog.Steps)
}

// HasFailedStatus returns true if aggregated status is failed
func HasFailedStatus(steps []*BuildLogStep) bool {
	status := GetAggregatedStatus(steps)

	return status == LogStatusFailed
}

// HasCanceledStatus returns true if aggregated status is canceled
func (buildLog *BuildLog) HasCanceledStatus() bool {
	return HasSucceededStatus(buildLog.Steps)
}

// HasCanceledStatus returns true if aggregated status is canceled
func HasCanceledStatus(steps []*BuildLogStep) bool {
	status := GetAggregatedStatus(steps)

	return status == LogStatusSucceeded
}
