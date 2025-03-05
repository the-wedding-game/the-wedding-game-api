package models

import (
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
)

var (
	submissions []Submission
)

func TestMain(m *testing.M) {
	submission1 := NewSubmission(1, 1, "answer1")
	submission2 := NewSubmission(1, 2, "answer2")
	submission3 := NewSubmission(1, 3, "answer3")
	submission4 := NewSubmission(1, 4, "answer4")
	submission5 := NewSubmission(1, 323, "answer4")

	submissions = []Submission{submission1, submission2, submission3, submission4, submission5}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestIsChallengeInSubmissionsWhenExists(t *testing.T) {
	assert.Equal(t, IsChallengeInSubmissions(1, submissions), true)
	assert.Equal(t, IsChallengeInSubmissions(2, submissions), true)
	assert.Equal(t, IsChallengeInSubmissions(3, submissions), true)
	assert.Equal(t, IsChallengeInSubmissions(4, submissions), true)
	assert.Equal(t, IsChallengeInSubmissions(323, submissions), true)
}

func TestIsChallengeInSubmissionsWhenNotExists(t *testing.T) {
	assert.Equal(t, IsChallengeInSubmissions(5, submissions), false)
	assert.Equal(t, IsChallengeInSubmissions(6, submissions), false)
	assert.Equal(t, IsChallengeInSubmissions(7, submissions), false)
	assert.Equal(t, IsChallengeInSubmissions(8, submissions), false)
	assert.Equal(t, IsChallengeInSubmissions(324, submissions), false)
}
