package utils

import (
	"gorm.io/gorm"
)

// PaginateResult 分页结果结构体
type PaginateResult struct {
	Total     int64       `json:"total"`     // 总记录数
	Page      int         `json:"page"`      // 当前页码
	PageSize  int         `json:"pageSize"`  // 每页大小
	TotalPage int64       `json:"totalPage"` // 总页数
	Data      interface{} `json:"data"`      // 数据
}

// Paginate GORM分页查询工具函数
// db: GORM数据库连接
// page: 页码（从1开始）
// pageSize: 每页大小
// result: 查询结果存储的切片指针
// 返回: 分页结果和错误信息
func Paginate(db *gorm.DB, page, pageSize int, result interface{}) (*PaginateResult, error) {
	// 设置默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 获取总记录数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 执行分页查询
	err = db.Offset(offset).Limit(pageSize).Find(result).Error
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPage := (total + int64(pageSize) - 1) / int64(pageSize)

	// 构建分页结果
	paginateResult := &PaginateResult{
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      result,
	}

	return paginateResult, nil
}

// PaginateWithCustomQuery 自定义查询的分页工具函数
// 适用于需要复杂查询条件的场景
func PaginateWithCustomQuery(countDB, queryDB *gorm.DB, page, pageSize int, result interface{}) (*PaginateResult, error) {
	// 设置默认值
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 获取总记录数（使用countDB）
	var total int64
	err := countDB.Count(&total).Error
	if err != nil {
		return nil, err
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 执行分页查询（使用queryDB）
	err = queryDB.Offset(offset).Limit(pageSize).Find(result).Error
	if err != nil {
		return nil, err
	}

	// 计算总页数
	totalPage := (total + int64(pageSize) - 1) / int64(pageSize)

	// 构建分页结果
	paginateResult := &PaginateResult{
		Total:     total,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
		Data:      result,
	}

	return paginateResult, nil
}

// BuildPaginateResponse 构建分页响应数据
// 用于API层返回统一格式的分页数据
func BuildPaginateResponse(data interface{}, total int64, page, pageSize int) map[string]interface{} {
	totalPages := (total + int64(pageSize) - 1) / int64(pageSize)
	return map[string]interface{}{
		"list":       data,
		"total":      total,
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": totalPages,
	}
}
