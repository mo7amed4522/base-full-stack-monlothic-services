package request

// AppCreateRequest 创建应用请求结构体
type AppCreateRequest struct {
	AppName      string `json:"appName" validate:"required"`       // 应用名称
	AppRemark    string `json:"appRemark"`                         // 应用备注
	AppIcon      string `json:"appIcon"`                           // 应用图标
	LanguageJson string `json:"language_json" validate:"required"` // 多语言配置JSON字符串
}

// AppManageSearchRequest 应用管理查询请求
type AppManageSearchRequest struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	AppName  string `json:"appName" form:"appName"`   // 应用名称
	Status   *int   `json:"status" form:"status"`     // 状态
}

// AppUpdateRequest 更新应用请求结构体
type AppUpdateRequest struct {
	ID           uint   `json:"id" validate:"required"`            // 记录ID
	AppName      string `json:"appName" validate:"required"`       // 应用名称
	AppRemark    string `json:"appRemark"`                         // 应用备注
	AppIcon      string `json:"appIcon"`                           // 应用图标
	LanguageJson string `json:"language_json" validate:"required"` // 多语言配置JSON字符串
}
