package utils

import (
	"github.com/go-playground/assert/v2"
	"os"
	"testing"
	"the-wedding-game-api/models"
)

var (
	submissions []models.Submission
)

func TestMain(m *testing.M) {
	submission1 := models.NewSubmission(1, 1, "answer1")
	submission2 := models.NewSubmission(1, 2, "answer2")
	submission3 := models.NewSubmission(1, 3, "answer3")
	submission4 := models.NewSubmission(1, 4, "answer4")
	submission5 := models.NewSubmission(1, 323, "answer4")

	submissions = []models.Submission{submission1, submission2, submission3, submission4, submission5}

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

func TestIsURLStrictWithValidUrls(t *testing.T) {
	assert.Equal(t, IsURLStrict("http://www.google.com"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com"), true)
	assert.Equal(t, IsURLStrict("http://google.com"), true)
	assert.Equal(t, IsURLStrict("https://google.com"), true)
	assert.Equal(t, IsURLStrict("http://www.google.com/"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com/"), true)
	assert.Equal(t, IsURLStrict("http://google.com/"), true)
	assert.Equal(t, IsURLStrict("https://google.com/"), true)
	assert.Equal(t, IsURLStrict("http://www.google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://www.google.com/path"), true)
	assert.Equal(t, IsURLStrict("http://google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://google.com/path"), true)
	assert.Equal(t, IsURLStrict("https://google.com/path/image.jpg"), true)
	assert.Equal(t, IsURLStrict("https://google.net/path/image.jpg"), true)
	assert.Equal(t, IsURLStrict("https://subdomain.google.net/path/image.jpg"), true)
}

func TestIsURLStrictWithInvalidUrls(t *testing.T) {
	assert.Equal(t, IsURLStrict("www.google.com"), false)
	assert.Equal(t, IsURLStrict("google.com"), false)
	assert.Equal(t, IsURLStrict("www.google.com/"), false)
	assert.Equal(t, IsURLStrict("google.com/"), false)
	assert.Equal(t, IsURLStrict("www.google.com/path"), false)
	assert.Equal(t, IsURLStrict("google.com/path"), false)
	assert.Equal(t, IsURLStrict("google.com/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("google.net/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("subdomain.google.net/path/image.jpg"), false)
	assert.Equal(t, IsURLStrict("http://"), false)
	assert.Equal(t, IsURLStrict("https://"), false)
	assert.Equal(t, IsURLStrict("https://invalid"), false)
	assert.Equal(t, IsURLStrict("www.google.com"), false)
}
