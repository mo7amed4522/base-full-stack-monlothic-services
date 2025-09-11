package request

// BxMatchesSearchRequest 比赛管理查询请求
type BxMatchesSearchRequest struct {
	Page       int    `json:"page" form:"page"`             // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
	Sport      string `json:"sport" form:"sport"`           // 运动类型: FOOTBALL, BASKETBALL
	Tsp        int64  `json:"tsp" form:"tsp"`               // 时间戳过滤
	MatchId    string `json:"matchId" form:"matchId"`       // 比赛ID
	Status     string `json:"status" form:"status"`         // 比赛状态: LIVE, UPCOMING, FINISHED
	IsHotMatch int    `json:"isHotMatch" form:"isHotMatch"` // 是否热门比赛
	IsTopMatch int    `json:"isTopMatch" form:"isTopMatch"` // 是否置顶比赛
}

// BatchMatchesRequest 批量操作比赛请求
type BatchMatchesRequest struct {
	MatchId []string `json:"matchId" binding:"required"` // 比赛ID列表
	Type    int      `json:"type" binding:"min=0,max=1"` // 1 确认 0 取消
}
