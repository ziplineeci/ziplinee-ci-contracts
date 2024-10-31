package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCiBuilderEvent(t *testing.T) {
	t.Run("ReturnsNoErrorWhenBuildIsSetForJobTypeBuild", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeBuild
		ciBuilderEvent.Build = &Build{}

		// act
		err := ciBuilderEvent.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenBuildIsNotSetForJobTypeBuild", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeBuild
		ciBuilderEvent.Build = nil

		// act
		err := ciBuilderEvent.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "build needs to be set for jobType build", err.Error())
	})

	t.Run("ReturnsNoErrorWhenReleaseIsSetForJobTypeRelease", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeRelease
		ciBuilderEvent.Release = &Release{}

		// act
		err := ciBuilderEvent.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenReleaseIsNotSetForJobTypeRelease", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeRelease
		ciBuilderEvent.Release = nil

		// act
		err := ciBuilderEvent.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "release needs to be set for jobType release", err.Error())
	})

	t.Run("ReturnsNoErrorWhenBotIsSetForJobTypeBot", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeBot
		ciBuilderEvent.Bot = &Bot{}

		// act
		err := ciBuilderEvent.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenBotIsNotSetForJobTypeBot", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.JobType = JobTypeBot
		ciBuilderEvent.Bot = nil

		// act
		err := ciBuilderEvent.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "bot needs to be set for jobType bot", err.Error())
	})

	t.Run("ReturnsNoErrorWhenGitIsSet", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.Git = &GitConfig{}

		// act
		err := ciBuilderEvent.Validate()

		assert.Nil(t, err)
	})

	t.Run("ReturnsErrorWhenGitIsNotSet", func(t *testing.T) {

		ciBuilderEvent := getCiBuilderEvent()
		ciBuilderEvent.Git = nil

		// act
		err := ciBuilderEvent.Validate()

		assert.NotNil(t, err)
		assert.Equal(t, "git needs to be set", err.Error())
	})
}

func getCiBuilderEvent() EstafetteCiBuilderEvent {
	return EstafetteCiBuilderEvent{
		JobType: JobTypeBot,
		JobName: "build-estafette-ci-api-12345",
		PodName: "build-estafette-ci-api-12345-abcd",
		Git:     &GitConfig{},
		Build:   &Build{},
		Release: &Release{},
		Bot:     &Bot{},
	}
}
