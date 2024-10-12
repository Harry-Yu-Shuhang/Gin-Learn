package config

const (
	RegisterErr = 4001 + iota //iota从0开始自增
	LoginErr
	GetPlayersErr
	GetRankingErr
	IncreaseVoteErr
)
