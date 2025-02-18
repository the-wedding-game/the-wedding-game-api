package utils

import "the-wedding-game-api/models"

func IsChallengeInSubmissions(challengeId uint, submissions []models.Submission) bool {
	for _, submission := range submissions {
		if submission.ChallengeID == challengeId {
			return true
		}
	}
	return false
}
