package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

const signkey = "sMDB3mOS00NDk3w8LTfg1"

func Sign(treeMap map[string]interface{}, signKey string) string {
	if signKey == "" {
		signKey = signkey
	}
	var sb strings.Builder

	keys := make([]string, 0, len(treeMap))
	for k := range treeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := treeMap[key]
		if value != nil {
			URLEncoded := IsURLEncoded(fmt.Sprintf("%v", value))
			if URLEncoded {
				value, _ = url.QueryUnescape(fmt.Sprintf("%v", value))
			}
			sb.WriteString(fmt.Sprintf("%s=%v&", key, value))
		}
	}

	sb.WriteString(fmt.Sprintf("key=%s", signKey))

	// 记录 MD5 前的明文
	plainText := sb.String()
	global.GVA_LOG.Info("MD5 前的明文: " + plainText)

	hasher := md5.New()
	hasher.Write([]byte(sb.String()))
	signature := hex.EncodeToString(hasher.Sum(nil))
	// 记录最终的签名
	global.GVA_LOG.Info("生成的签名: " + strings.ToUpper(signature))

	return strings.ToUpper(signature)
}

func IsURLEncoded(s string) bool {
	// 尝试解码
	decoded, err := url.QueryUnescape(s)
	if err != nil {
		// 解码失败，可能是未编码的字符串
		return false
	}

	// 比较解码前后的字符串
	if decoded != s {
		// 解码后的字符串与原字符串不同，说明原字符串可能是经过 URL 编码的
		return true
	}

	// 如果解码后的字符串与原字符串相同，可能未编码
	return false
}

// 支付订单和出款订单接口专用签名
func OrderSign(treeMap map[string]interface{}, signKey string) string {
	if signKey == "" {
		signKey = signkey
	}
	var sb strings.Builder

	keys := make([]string, 0, len(treeMap))
	for k := range treeMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := treeMap[key]
		if value != nil && value != "" {
			URLEncoded := IsURLEncoded(fmt.Sprintf("%v", value))
			if URLEncoded {
				value, _ = url.QueryUnescape(fmt.Sprintf("%v", value))
			}
			sb.WriteString(fmt.Sprintf("%s=%v&", key, value))
		}
	}

	sb.WriteString(fmt.Sprintf("signKey=%s", signKey))

	// 记录 MD5 前的明文
	plainText := sb.String()
	global.GVA_LOG.Info("MD5 前的明文: " + plainText)

	hasher := md5.New()
	hasher.Write([]byte(sb.String()))
	signature := hex.EncodeToString(hasher.Sum(nil))
	// 记录最终的签名
	global.GVA_LOG.Info("生成的签名: " + signature)

	return signature
}

func StructToMapNew(obj interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	objValue := reflect.ValueOf(obj).Elem()
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldInfo := objType.Field(i)
		mapKey := fieldInfo.Name

		// 获取 json 标签，如果存在则使用标签值
		if tag := fieldInfo.Tag.Get("json"); tag != "" {
			index := strings.Index(tag, ",")
			if index != -1 {
				// 截取逗号前的部分
				tag = tag[:index]
			}
			mapKey = tag
		}

		// 将字段名称的首字母转换为小写
		mapKey = toLowerFirst(mapKey)
		result[mapKey] = field.Interface()
	}

	return result, nil
}

// 将字符串的首字母转换为小写
func toLowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
