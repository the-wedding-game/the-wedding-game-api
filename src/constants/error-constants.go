package constants

import (
	"fmt"
	"the-wedding-game-api/config"
)

var InvalidChallengeTypeError = "invalid challenge type"
var AnswerRequiredError = "answer is required for answer question challenges"
var InvalidImageURLError = "invalid image url"
var InvalidChallengeIDError = "invalid challenge id"
var ImageIsRequiredError = "image is required"
var FileMustBeAnImageError = "file must be an image"
var FileIsEmptyError = "file is empty"
var MaxFileSizeError = fmt.Sprintf("maximum file size is %d bytes", config.MAX_UPLOAD_SIZE) // 10 MB
