package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/bxapi"

// TeamInfo 队伍信息
type TeamInfo struct {
	Id    string `json:"id"`
	Logo  string `json:"logo"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

// CompetitionInfo 赛事信息
type CompetitionInfo struct {
	Id   string `json:"id"`
	Logo string `json:"logo"`
	Name string `json:"name"`
}

// CBxMatchSimpleResponse 简单比赛响应（包含联赛名称）
type CBxMatchSimpleResponse struct {
	MatchId         string                  `json:"matchId"`         // 比赛ID
	HomeTeam        TeamInfo                `json:"homeTeam"`        // 主队信息
	AwayTeam        TeamInfo                `json:"awayTeam"`        // 客队信息
	MatchStartTime  int                     `json:"matchStartTime"`  // 比赛开始时间
	CurrentGameTime int                     `json:"currentGameTime"` // 当前比赛时间
	MatchStatus     string                  `json:"matchStatus"`     // 比赛状态
	CoinsPrice      int                     `json:"CoinsPrice"`      // 金币
	Status          string                  `json:"status"`          // 状态描述
	HasTracker      bool                    `json:"hasTracker"`      // 是否有跟踪器
	HasLineup       bool                    `json:"hasLineup"`       // 是否有阵容
	IsHotMatch      bool                    `json:"isHotMatch"`      // 是否热门比赛
	IsTopMatch      bool                    `json:"isTopMatch"`      // 是否置顶比赛
	HasStream       bool                    `json:"hasStream"`       // 是否有直播
	StreamId        string                  `json:"streamId"`        // 是否有直播
	LeagueName      string                  `json:"leagueName"`      // 联赛名称
	LeagueLogo      string                  `json:"leagueLogo"`      // 联赛Logo
	LeagueId        string                  `json:"leagueId"`        // 联赛ID
	HasMarkets      bool                    `json:"hasMarkets"`      // 是否有投注市场
	Markets         []bxapi.BxBettingMarket `json:"markets"`         // 投注市场
	VoteEnabled     bool                    `json:"voteEnabled"`     // 是否投票
}

// CBxMatchSimpleListResponse 简单比赛列表响应
type CBxMatchSimpleListResponse struct {
	Matches []CBxMatchSimpleResponse `json:"matches"` // 比赛列表
	Total   int64                    `json:"total"`   // 总数
	Page    int                      `json:"page"`    // 当前页
	Size    int                      `json:"size"`    // 每页大小
}
