package contracts

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildLog(t *testing.T) {
	t.Run("JSONMarshalSingleBuildLog", func(t *testing.T) {
		buildLog := BuildLog{
			ID:           "5",
			RepoSource:   "github.com",
			RepoOwner:    "estafette",
			RepoName:     "estafette-ci-api",
			RepoBranch:   "master",
			RepoRevision: "as23456",
			Steps: []*BuildLogStep{
				&BuildLogStep{
					Step: "init",
					Image: &BuildLogStepDockerImage{
						Name:         "golang",
						Tag:          "1.10.2-alpine3.7",
						IsPulled:     false,
						ImageSize:    135000,
						PullDuration: 2 * time.Second,
						Error:        "",
					},
					Duration: 91 * time.Second,
					LogLines: []BuildLogLine{
						BuildLogLine{
							Timestamp:  time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
							StreamType: "stdout",
							Text: "ok  	github.com/estafette/estafette-ci-contracts	0.017s",
						},
					},
					ExitCode: 0,
					Status:   LogStatusSucceeded,
				},
			},
			InsertedAt: time.Date(2018, 4, 17, 8, 3, 0, 0, time.UTC),
		}
		// act
		bytes, err := json.Marshal(&buildLog)
		assert.Nil(t, err)
		assert.Equal(t, "{\"id\":\"5\",\"repoSource\":\"github.com\",\"repoOwner\":\"estafette\",\"repoName\":\"estafette-ci-api\",\"repoBranch\":\"master\",\"repoRevision\":\"as23456\",\"buildID\":\"\",\"steps\":[{\"step\":\"init\",\"image\":{\"name\":\"golang\",\"tag\":\"1.10.2-alpine3.7\",\"isPulled\":false,\"imageSize\":135000,\"pullDuration\":2000000000},\"duration\":91000000000,\"logLines\":[{\"timestamp\":\"2018-04-17T08:03:00Z\",\"streamType\":\"stdout\",\"text\":\"ok  \\tgithub.com/estafette/estafette-ci-contracts\\t0.017s\"}],\"exitCode\":0,\"status\":\"SUCCEEDED\"}],\"insertedAt\":\"2018-04-17T08:03:00Z\"}", string(bytes))
	})
}

func TestHasSucceededStatus(t *testing.T) {
	t.Run("ReturnsFalseIfNoSteps", func(t *testing.T) {

		steps := []*BuildLogStep{}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.False(t, succeeded)
	})

	t.Run("ReturnsFalseIfAllStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.False(t, succeeded)
	})

	t.Run("ReturnsFalseIfAnyStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.False(t, succeeded)
	})

	t.Run("ReturnsFalseIfAnyStepsCanceled", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusCanceled,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusCanceled,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.False(t, succeeded)
	})

	t.Run("ReturnsTrueIfAStepFailedButSucceededInRetry", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.True(t, succeeded)
	})

	t.Run("ReturnsTrueIfSomeStepsAreSkipped", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "SKIPPED",
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.True(t, succeeded)
	})

	t.Run("ReturnsFalseIfAStepFailedButSucceededInRetryButAnotherStepFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusFailed,
			},
		}

		// act
		succeeded := HasSucceededStatus(steps)

		assert.False(t, succeeded)
	})
}

func TestGetAggregatedStatus(t *testing.T) {
	t.Run("ReturnsUnknownIfNoSteps", func(t *testing.T) {

		steps := []*BuildLogStep{}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusUnknown, status)
	})

	t.Run("ReturnsSucceededIfAllStepsSucceeded", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusSucceeded,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusSucceeded, status)
	})

	t.Run("ReturnsFailedIfAllStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusFailed, status)
	})

	t.Run("ReturnsFailedIfAnyStepsFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusFailed, status)
	})

	t.Run("ReturnsCanceledIfAnyStepsCanceled", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusCanceled,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusCanceled, status)
	})

	t.Run("ReturnsSucceededIfAStepFailedButSucceededInRetry", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusSucceeded, status)
	})

	t.Run("ReturnsSucceededIfSomeStepsAreSkipped", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: "SKIPPED",
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusSucceeded,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusSucceeded, status)
	})

	t.Run("ReturnsFailedIfAStepFailedButSucceededInRetryButAnotherStepFailed", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusFailed,
			},
			&BuildLogStep{
				Step:     "stage-b",
				RunIndex: 1,
				Status:   LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusFailed,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusFailed, status)
	})

	t.Run("ReturnsCanceledIfAnyStepsCanceled", func(t *testing.T) {

		steps := []*BuildLogStep{
			&BuildLogStep{
				Step:   "stage-a",
				Status: LogStatusSucceeded,
			},
			&BuildLogStep{
				Step:   "stage-b",
				Status: LogStatusCanceled,
			},
			&BuildLogStep{
				Step:   "stage-c",
				Status: LogStatusCanceled,
			},
		}

		// act
		status := GetAggregatedStatus(steps)

		assert.Equal(t, LogStatusCanceled, status)
	})
}
