package utils

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bxapi"
)

// DictUtils 字典工具类
type DictUtils struct{}

var Dict = new(DictUtils)

// GetDictValue 根据字典类型获取第一个值
// @param dictType 字典类型（使用constants中定义的常量）
// @return string 字典值
func (d *DictUtils) GetDictValue(dictType string) (string, error) {
	if global.GVA_DB == nil {
		return "", fmt.Errorf("database not initialized")
	}

	var dict bxapi.BxDict
	err := global.GVA_DB.Where("type = ? AND is_deleted = ?", dictType, false).
		Order("sort ASC, id ASC").
		First(&dict).Error

	if err != nil {
		return "", fmt.Errorf("failed to get dict value: %v", err)
	}

	return dict.Value, nil
}

// GetDictValues 批量获取字典值
// @param dictTypes 字典类型数组
// @return map[string]string 字典类型到值的映射
func (d *DictUtils) GetDictValues(dictTypes []string) (map[string]string, error) {
	if global.GVA_DB == nil {
		return nil, fmt.Errorf("database not initialized")
	}

	var dicts []bxapi.BxDict
	err := global.GVA_DB.Where("type IN ? AND is_deleted = ?", dictTypes, false).
		Order("sort ASC, id ASC").
		Find(&dicts).Error

	if err != nil {
		return nil, fmt.Errorf("failed to get dict values: %v", err)
	}

	result := make(map[string]string)
	for _, dict := range dicts {
		result[dict.Type] = dict.Value
	}

	// 确保所有请求的类型都有返回值，缺失的设为空字符串
	for _, dictType := range dictTypes {
		if _, exists := result[dictType]; !exists {
			result[dictType] = ""
		}
	}

	return result, nil
}
