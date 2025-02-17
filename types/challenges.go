package types

type ChallengeType string

const (
	UploadPhotoChallenge    ChallengeType = "UPLOAD_PHOTO"
	AnswerQuestionChallenge ChallengeType = "ANSWER_QUESTION"
)

type ChallengeStatus string

const (
	ActiveChallenge   ChallengeStatus = "ACTIVE"
	InactiveChallenge ChallengeStatus = "INACTIVE"
)

type CreateChallengeRequest struct {
	Name        string        `json:"name" binding:"required" validate:"required"`
	Description string        `json:"description" binding:"required" validate:"required"`
	Points      uint          `json:"points" binding:"required" validate:"required,gte=0"`
	Image       string        `json:"image" binding:"required" validate:"required,url"`
	Type        ChallengeType `json:"type" binding:"required" validate:"required"`
	Answer      string        `json:"answer"`
}

type ChallengeCreatedResponse struct {
	Id          uint            `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Points      uint            `json:"points"`
	Image       string          `json:"image"`
	Status      ChallengeStatus `json:"status"`
	Type        ChallengeType   `json:"type"`
}

type GetChallengesResponse struct {
	Challenges []ChallengeCreatedResponse `json:"challenges"`
}
