package contracts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogStatusEquals(t *testing.T) {
	t.Run("ReturnsTrueWhenBothArePending", func(t *testing.T) {
		// act
		equal := LogStatusPending.Equals(StatusPending)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothArePending", func(t *testing.T) {
		// act
		equal := LogStatusPending.Equals(StatusRunning)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreRunning", func(t *testing.T) {
		// act
		equal := LogStatusRunning.Equals(StatusRunning)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreRunning", func(t *testing.T) {
		// act
		equal := LogStatusRunning.Equals(StatusCanceled)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreSucceeded", func(t *testing.T) {
		// act
		equal := LogStatusSucceeded.Equals(StatusSucceeded)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreSucceeded", func(t *testing.T) {
		// act
		equal := LogStatusSucceeded.Equals(StatusFailed)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreFailed", func(t *testing.T) {
		// act
		equal := LogStatusFailed.Equals(StatusFailed)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreFailed", func(t *testing.T) {
		// act
		equal := LogStatusFailed.Equals(StatusSucceeded)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreCanceled", func(t *testing.T) {
		// act
		equal := LogStatusCanceled.Equals(StatusCanceled)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreCanceled", func(t *testing.T) {
		// act
		equal := LogStatusCanceled.Equals(StatusSucceeded)

		assert.False(t, equal)
	})

	t.Run("ReturnsFalseWhenBothAreUnknown", func(t *testing.T) {
		// act
		equal := LogStatusUnknown.Equals(StatusUnknown)

		assert.False(t, equal)
	})
}

func TestStatusEquals(t *testing.T) {
	t.Run("ReturnsTrueWhenBothArePending", func(t *testing.T) {
		// act
		equal := StatusPending.Equals(LogStatusPending)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothArePending", func(t *testing.T) {
		// act
		equal := StatusPending.Equals(LogStatusRunning)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreRunning", func(t *testing.T) {
		// act
		equal := StatusRunning.Equals(LogStatusRunning)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreRunning", func(t *testing.T) {
		// act
		equal := StatusRunning.Equals(LogStatusCanceled)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreSucceeded", func(t *testing.T) {
		// act
		equal := StatusSucceeded.Equals(LogStatusSucceeded)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreSucceeded", func(t *testing.T) {
		// act
		equal := StatusSucceeded.Equals(LogStatusFailed)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreFailed", func(t *testing.T) {
		// act
		equal := StatusFailed.Equals(LogStatusFailed)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreFailed", func(t *testing.T) {
		// act
		equal := StatusFailed.Equals(LogStatusSucceeded)

		assert.False(t, equal)
	})

	t.Run("ReturnsTrueWhenBothAreCanceled", func(t *testing.T) {
		// act
		equal := StatusCanceled.Equals(LogStatusCanceled)

		assert.True(t, equal)
	})

	t.Run("ReturnsFalseWhenNotBothAreCanceled", func(t *testing.T) {
		// act
		equal := StatusCanceled.Equals(LogStatusSucceeded)

		assert.False(t, equal)
	})

	t.Run("ReturnsFalseWhenBothAreUnknown", func(t *testing.T) {
		// act
		equal := StatusUnknown.Equals(LogStatusUnknown)

		assert.False(t, equal)
	})
}
