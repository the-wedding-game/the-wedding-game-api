package types

type CurrentUserPointsResponse struct {
	Points uint `json:"points"`
}

type LeaderboardEntry struct {
	Username string `json:"username"`
	Points   uint   `json:"points"`
}

type GetLeaderboardResponse struct {
	Leaderboard []LeaderboardEntry `json:"leaderboard"`
}
