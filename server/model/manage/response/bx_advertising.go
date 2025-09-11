package response

import (
	"time"
)

// BxAdvertisingResponse 广告响应
type BxAdvertisingResponse struct {
	AdId          uint64    `json:"ad_id"`           // 广告ID
	SiteId        uint      `json:"site_id"`         // 站点ID
	Title         string    `json:"title"`           // 标题
	Advertiser    string    `json:"advertiser"`      // 广告商
	AndroidUrl    string    `json:"android_url"`     // 安卓跳转链接
	IosUrl        string    `json:"ios_url"`         // iOS跳转链接
	H5Url         string    `json:"h5_url"`          // H5跳转链接
	ImgUrl        string    `json:"img_url"`         // 图片地址
	CornerMarkUrl string    `json:"corner_mark_url"` // 角标URL
	Description   string    `json:"description"`     // 描述
	JumpType      uint8     `json:"jump_type"`       // 跳转类型
	ClickCount    uint      `json:"click_count"`     // 点击数量
	Sort          uint      `json:"sort"`            // 排序
	Status        uint8     `json:"status"`          // 状态
	CreatedAt     time.Time `json:"created_at"`      // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`      // 更新时间
	Positions     []int     `json:"positions"`       // 关联的位置数组
	PageType      int       `json:"page_type"`       // 页面类型
}
