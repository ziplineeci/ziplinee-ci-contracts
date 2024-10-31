package contracts

import "strings"

type Status string

const (
	// StatusPending indicates container is pulling
	StatusPending Status = "pending"
	// StatusRunning indicates container is running
	StatusRunning Status = "running"
	// StatusSucceeded indicates execution was successful
	StatusSucceeded Status = "succeeded"
	// StatusFailed indicates execution was not successful
	StatusFailed Status = "failed"
	// StatusCanceling indicates execution is canceling
	StatusCanceling Status = "canceling"
	// StatusCanceled indicates execution was canceled
	StatusCanceled Status = "canceled"

	// StatusUnknown provides a default but not allowed status for unmarshalling
	StatusUnknown Status = ""
)

type LogStatus string

const (
	// LogStatusUnknown indicates execution never started for some reason
	LogStatusUnknown LogStatus = "UNKNOWN"
	// LogStatusSucceeded indicates execution was successful
	LogStatusSucceeded LogStatus = "SUCCEEDED"
	// LogStatusFailed indicates execution was not successful
	LogStatusFailed LogStatus = "FAILED"
	// LogStatusSkipped indicates execution was skipped
	LogStatusSkipped LogStatus = "SKIPPED"
	// LogStatusCanceled indicates execution was canceled
	LogStatusCanceled LogStatus = "CANCELED"
	// LogStatusPending indicates container is pulling
	LogStatusPending LogStatus = "PENDING"
	// LogStatusRunning indicates container is running
	LogStatusRunning LogStatus = "RUNNING"
)

type LogType string

const (
	// TypeStage indicates that a tail message is for a main stage or parallel stage
	LogTypeStage LogType = "stage"
	// TypeService indicates that a tail message is for a service container
	LogTypeService LogType = "service"
)

func (l LogStatus) Equals(s Status) bool {
	return strings.ToLower(string(l)) == strings.ToLower(string(s))
}

func (s Status) Equals(l LogStatus) bool {
	return l.Equals(s)
}

func (l LogStatus) ToStatus() Status {
	switch l {
	case LogStatusSucceeded:
		return StatusSucceeded
	case LogStatusFailed:
		return StatusFailed
	case LogStatusCanceled:
		return StatusCanceled
	case LogStatusPending:
		return StatusPending
	case LogStatusRunning:
		return StatusRunning
	}

	return StatusUnknown
}

func (s Status) ToLogStatus() LogStatus {
	switch s {
	case StatusPending:
		return LogStatusPending
	case StatusRunning:
		return LogStatusRunning
	case StatusSucceeded:
		return LogStatusSucceeded
	case StatusFailed:
		return LogStatusFailed
	case StatusCanceled:
		return LogStatusCanceled
	}

	return LogStatusUnknown
}
