package constants

// 奖励类型常量 - Reward Type Constants
const (
	REWARD_ADD_FAVORITE_MATCH   = 1001 // 添加收藏比赛 - Add favorite match
	REWARD_COMMENT_LIVE_CHAT    = 1002 // 直播聊天评论 - Comment in live chat
	REWARD_DAILY_LOGIN          = 1003 // 每日登录 - Daily login
	REWARD_INVITE_FRIEND        = 1004 // 邀请好友 - Invite friend
	REWARD_JOIN_TELEGRAM        = 1005 // 加入Telegram群组 - Join Telegram group
	REWARD_RATE_APP             = 1006 // 为应用评分 - Rate app
	REWARD_REGISTER_ACCOUNT     = 1007 // 注册账户 - Register account
	REWARD_WATCH_LIVE_STREAM    = 1008 // 观看直播流 - Watch live stream
	REWARD_WATCH_LIVE_TRACKER   = 1009 // 观看直播追踪器 - Watch live tracker
	REWARD_DAILY_LOGIN_STREAK_7 = 1010 // 7天连续签到奖励 - 7-day login streak bonus
)

// 收藏类型常量 - Favorite target types
const (
	FAVORITE_TYPE_MATCH  = 1 // 比赛类型 - Match type
	FAVORITE_TYPE_TEAM   = 2 // 球队类型 - Team type
	FAVORITE_TYPE_LEAGUE = 3 // 联赛类型 - League type
)

// 字典类型常量 - Dictionary Type Constants
const (
	DICT_TYPE_ANDROID_AD_ENABLE = "ANDROID_AD_ENABLE" // android 广告是否启用 - Advertisement enable status
	DICT_TYPE_IOS_AD_ENABLE     = "IOS_AD_ENABLE"     // ios 广告是否启用 - Advertisement enable status
	DICT_TYPE_H5_AD_ENABLE      = "H5_AD_ENABLE"      // h5 广告是否启用 - H5 advertisement enable status
	DICT_TYPE_AD_ENABLE         = "AD_ENABLE"         // 广告是否启用 - General advertisement enable status
)
