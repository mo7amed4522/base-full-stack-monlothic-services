package request

// BxAdvertisingSearchRequest 广告搜索请求
type BxAdvertisingSearchRequest struct {
	Page       int    `json:"page" form:"page"`             // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
	AdId       string `json:"ad_id" form:"ad_id"`           // 广告ID搜索
	Title      string `json:"title" form:"title"`           // 标题搜索
	Advertiser string `json:"advertiser" form:"advertiser"` // 广告商搜索
	SiteId     string `json:"site_id" form:"site_id"`       // 站点筛选
	Status     string `json:"status" form:"status"`         // 状态筛选
	PageType   string `json:"page_type" form:"page_type"`   // 页面类型筛选
	StartTime  string `json:"startTime" form:"startTime"`   // 开始时间
	EndTime    string `json:"endTime" form:"endTime"`       // 结束时间
}

// BxAdvertisingCreateRequest 创建广告请求
type BxAdvertisingCreateRequest struct {
	SiteId        uint    `json:"site_id"`                                  // 站点ID
	Title         string  `json:"title" binding:"required"`                 // 标题
	Advertiser    string  `json:"advertiser"`                               // 广告商
	AndroidUrl    string  `json:"android_url"`                              // 安卓跳转链接
	IosUrl        string  `json:"ios_url"`                                  // iOS跳转链接
	H5Url         string  `json:"h5_url"`                                   // H5跳转链接
	ImgUrl        string  `json:"img_url"`                                  // 图片地址
	CornerMarkUrl string  `json:"corner_mark_url"`                          // 角标URL
	Description   string  `json:"description"`                              // 描述
	JumpType      uint8   `json:"jump_type" binding:"required,min=1,max=2"` // 跳转类型
	Position      []uint8 `json:"position" binding:"required"`              // 位置数组
	PageType      uint8   `json:"page_type"`                                // 页面类型
	Sort          uint    `json:"sort"`                                     // 排序
	Status        uint8   `json:"status" binding:"min=0,max=1"`             // 状态
}

// BxAdvertisingUpdateRequest 更新广告请求
type BxAdvertisingUpdateRequest struct {
	AdId          uint64  `json:"ad_id" binding:"required"`                 // 广告ID
	SiteId        uint    `json:"site_id"`                                  // 站点ID
	Title         string  `json:"title" binding:"required"`                 // 标题
	Advertiser    string  `json:"advertiser"`                               // 广告商
	AndroidUrl    string  `json:"android_url"`                              // 安卓跳转链接
	IosUrl        string  `json:"ios_url"`                                  // iOS跳转链接
	H5Url         string  `json:"h5_url"`                                   // H5跳转链接
	ImgUrl        string  `json:"img_url"`                                  // 图片地址
	CornerMarkUrl string  `json:"corner_mark_url"`                          // 角标URL
	Description   string  `json:"description"`                              // 描述
	JumpType      uint8   `json:"jump_type" binding:"required,min=1,max=2"` // 跳转类型
	Position      []uint8 `json:"position" binding:"required"`              // 位置数组
	PageType      uint8   `json:"page_type"`                                // 页面类型
	Sort          uint    `json:"sort"`                                     // 排序
	Status        uint8   `json:"status" binding:"min=0,max=1"`             // 状态
}

// BxAdvertisingDeleteRequest 删除广告请求
type BxAdvertisingDeleteRequest struct {
	AdId uint64 `json:"ad_id" form:"ad_id" binding:"required"` // 广告ID
}

// BxDictUpdateRequest 更新字典请求
type BxDictUpdateRequest struct {
	Type  string `json:"type" binding:"required"`  // 字典类型
	Value string `json:"value" binding:"required"` // 字典值
}
