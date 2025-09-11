package manage

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/model/bxapi"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/response"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/manage/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type CBxMatchesService struct{}

// GetBxMatchesList 获取比赛列表
func (cBxMatchesService *CBxMatchesService) GetBxMatchesList(req request.BxMatchesSearchRequest) (matchesResponse []response.CBxMatchSimpleResponse, total int64, err error) {
	if global.GVA_DB == nil {
		return nil, 0, fmt.Errorf("db not init")
	}

	// 设置默认分页参数
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 构建查询条件
	query := global.GVA_DB.Model(&bxapi.BxMatches{})

	// 运动类型过滤
	if req.Sport != "" {
		query = query.Where("sport = ?", strings.ToUpper(req.Sport))
	}

	// 时间戳过滤 - 获取指定时间之后的比赛
	if req.Tsp > 0 {
		query = query.Where("match_start_time >= ?", req.Tsp)
	}

	// 比赛ID过滤
	if req.MatchId != "" {
		query = query.Where("match_id LIKE ?", "%"+req.MatchId+"%")
	}

	// 比赛状态过滤
	if req.Status != "" {
		query = query.Where("match_status = ?", strings.ToUpper(req.Status))
	}

	// 热门比赛过滤
	if req.IsHotMatch != 0 {
		query = query.Where("is_hot_match = ?", req.IsHotMatch)
	}

	// 置顶比赛过滤
	if req.IsTopMatch != 0 {
		query = query.Where("is_top_match = ?", req.IsTopMatch)
	}

	// 使用分页工具进行查询
	var matches []bxapi.BxMatches
	paginateResult, err := utils.PaginateWithCustomQuery(
		query,                                // count查询
		query.Order("match_start_time DESC"), // 数据查询，按比赛开始时间倒序
		req.Page, req.PageSize, &matches,
	)
	if err != nil {
		return matchesResponse, 0, err
	}

	matchesResponse, err = cBxMatchesService.buildFromQueryResults(matches)

	return matchesResponse, paginateResult.Total, nil
}

// buildFromQueryResults 从查询结果构建响应
func (cBxMatchesService *CBxMatchesService) buildFromQueryResults(results []bxapi.BxMatches) (matchesResponse []response.CBxMatchSimpleResponse, err error) {
	for _, result := range results {
		// 解析Competition JSON
		var competition response.CompetitionInfo
		if err := json.Unmarshal([]byte(result.Competition), &competition); err != nil {
			continue // 跳过解析失败的比赛
		}

		// 解析HomeTeam JSON
		var homeTeam response.TeamInfo
		if err := json.Unmarshal([]byte(result.HomeTeam), &homeTeam); err != nil {
			continue
		}

		// 解析AwayTeam JSON
		var awayTeam response.TeamInfo
		if err := json.Unmarshal([]byte(result.AwayTeam), &awayTeam); err != nil {
			continue
		}

		// 查询该比赛的投注市场信息
		var markets []bxapi.BxBettingMarket
		err = global.GVA_DB.Where("match_id = ?", result.MatchId).Find(&markets).Error
		if err != nil {
			// 如果查询投注市场失败，继续处理，不影响比赛信息
			markets = []bxapi.BxBettingMarket{}
		}

		// 构建简单比赛信息
		matchInfo := response.CBxMatchSimpleResponse{
			MatchId:        result.MatchId,
			HomeTeam:       homeTeam,
			AwayTeam:       awayTeam,
			MatchStartTime: result.MatchStartTime,
			MatchStatus:    result.MatchStatus,
			Status:         result.Status,
			LeagueName:     competition.Name,
			LeagueLogo:     competition.Logo,
			LeagueId:       competition.Id,
			Markets:        markets,
			IsHotMatch:     result.IsHotMatch,
			IsTopMatch:     result.IsTopMatch,
			HasTracker:     result.HasTracker,
			HasLineup:      result.HasLineups,
			HasStream:      result.HasStream,
			StreamId:       result.StreamId,
			VoteEnabled:    result.VoteEnabled,
		}

		matchesResponse = append(matchesResponse, matchInfo)
	}
	return matchesResponse, nil
}

// BatchSetHot 批量设置热门
func (cBxMatchesService *CBxMatchesService) BatchSetHot(matchesRequest request.BatchMatchesRequest) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}
	matchIds := matchesRequest.MatchId
	if len(matchIds) == 0 {
		return fmt.Errorf("比赛ID列表不能为空")
	}

	err := global.GVA_DB.Model(&bxapi.BxMatches{}).
		Where("match_id IN ?", matchIds).
		Update("is_hot_match", matchesRequest.Type).Error

	if err != nil {
		return fmt.Errorf("批量设置热门失败: %v", err)
	}

	return nil
}

// BatchSetTop 批量置顶
func (cBxMatchesService *CBxMatchesService) BatchSetTop(matchesRequest request.BatchMatchesRequest) error {
	if global.GVA_DB == nil {
		return fmt.Errorf("db not init")
	}

	matchIds := matchesRequest.MatchId
	if len(matchIds) == 0 {
		return fmt.Errorf("比赛ID列表不能为空")
	}

	err := global.GVA_DB.Model(&bxapi.BxMatches{}).
		Where("match_id IN ?", matchIds).
		Update("is_top_match", matchesRequest.Type).Error

	if err != nil {
		return fmt.Errorf("批量置顶失败: %v", err)
	}

	return nil
}
