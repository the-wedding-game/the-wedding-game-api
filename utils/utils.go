package utils

import (
	"net/url"
	"regexp"
	"strings"
	"the-wedding-game-api/models"
)

func IsChallengeInSubmissions(challengeId uint, submissions []models.Submission) bool {
	for _, submission := range submissions {
		if submission.ChallengeID == challengeId {
			return true
		}
	}
	return false
}

func IsURLStrict(s string) bool {
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	if u.Scheme == "" || u.Host == "" { // Must have scheme and host
		return false
	}

	if strings.Contains(u.Host, "localhost") {
		r := regexp.MustCompile(`^(?:http(s)?://)?localhost(?::\d+)?(?:/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`)
		return r.MatchString(s)
	}

	r := regexp.MustCompile(`^(?:http(s)?://)?[\w.-]+(?:\.[\w.-]+)+(?::\d+)?(?:/[\w\-._~:/?#[\]@!$&'()*+,;=]*)?$`)
	return r.MatchString(s)
}
