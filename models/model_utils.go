package models

func IsChallengeInSubmissions(challengeId uint, submissions []Submission) bool {
	for _, submission := range submissions {
		if submission.ChallengeID == challengeId {
			return true
		}
	}
	return false
}
