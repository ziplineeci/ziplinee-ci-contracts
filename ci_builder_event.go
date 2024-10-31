package contracts

import (
	"errors"
)

type BuildEventType string

const (
	BuildEventTypeUnknown      BuildEventType = ""
	BuildEventTypeUpdateStatus BuildEventType = "updateStatus"
	BuildEventTypeClean        BuildEventType = "clean"
)

type EstafetteCiBuilderEvent struct {
	BuildEventType BuildEventType `json:"buildEventType,omitempty"`
	JobType        JobType        `json:"jobType,omitempty"`
	JobName        string         `json:"job_name"`
	PodName        string         `json:"pod_name,omitempty"`

	Build   *Build     `json:"build,omitempty"`
	Release *Release   `json:"release,omitempty"`
	Bot     *Bot       `json:"bot,omitempty"`
	Git     *GitConfig `json:"git,omitempty"`
}

func (bc *EstafetteCiBuilderEvent) Validate() (err error) {

	if bc.Git == nil {
		return errors.New("git needs to be set")
	}

	switch bc.JobType {
	case JobTypeBuild:
		if bc.Build == nil {
			return errors.New("build needs to be set for jobType build")
		}
	case JobTypeRelease:
		if bc.Release == nil {
			return errors.New("release needs to be set for jobType release")
		}
	case JobTypeBot:
		if bc.Bot == nil {
			return errors.New("bot needs to be set for jobType bot")
		}
	}

	return nil
}

func (bc *EstafetteCiBuilderEvent) GetStatus() Status {
	switch bc.JobType {
	case JobTypeBuild:
		if bc.Build != nil {
			return bc.Build.BuildStatus
		}
	case JobTypeRelease:
		if bc.Release != nil {
			return bc.Release.ReleaseStatus
		}
	case JobTypeBot:
		if bc.Bot != nil {
			return bc.Bot.BotStatus
		}
	}

	return StatusUnknown
}

func (bc *EstafetteCiBuilderEvent) SetStatus(status Status) {
	switch bc.JobType {
	case JobTypeBuild:
		if bc.Build != nil {
			bc.Build.BuildStatus = status
		}
	case JobTypeRelease:
		if bc.Release != nil {
			bc.Release.ReleaseStatus = status
		}
	case JobTypeBot:
		if bc.Bot != nil {
			bc.Bot.BotStatus = status
		}
	}
}
