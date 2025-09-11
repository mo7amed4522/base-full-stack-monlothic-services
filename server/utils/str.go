package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func StringToInt64(str string) int64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return int64(i)
}

// 将 []string 转换为 []int 的函数，跳过无效的字符串
func StrToIntSlice(strSlice []string) []int {
	var intSlice []int

	for _, str := range strSlice {
		intValue, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		intSlice = append(intSlice, intValue)
	}

	return intSlice
}

// RandString 自动生成一个12为的随机字符串
func RandString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// FormatAmount formats the amount to two decimal places if necessary
func FormatAmount(amount float64) string {
	// Convert float64 to string with 2 decimal places
	amountStr := fmt.Sprintf("%.2f", amount)

	// Check if the amount has exactly two decimal places
	if strings.Contains(amountStr, ".") {
		// Split the amount into integer and decimal parts
		parts := strings.Split(amountStr, ".")
		if len(parts) == 2 {
			// If decimal part is zero, ensure .00 is shown
			if parts[1] == "00" {
				return parts[0] + ".00"
			}
			// If the decimal part is less than 10 (e.g., 100.2), ensure it has two digits
			if len(parts[1]) == 1 {
				parts[1] = parts[1] + "0"
			}
			return parts[0] + "." + parts[1]
		}
	}

	return amountStr
}

// IsTwoDecimalPlaces 检查浮点数金额是否精确到小数点后两位。
// 这个函数通过将金额转换成字符串，然后检查字符串的小数部分长度是否为2来实现。
// 参数:
//
//	amount - 需要检查的浮点数金额。
//
// 返回值:
//
//	如果金额精确到小数点后两位，则返回true；否则返回false。
func IsValidAmount(amount string) bool {
	// 使用正则表达式匹配格式：数字，后跟小数点，后面是两位数字
	regex := `^\d+\.\d{2}$`
	matched, _ := regexp.MatchString(regex, amount)
	return matched
}

func GetStartTime(startTime string) (t time.Time, err error) {
	//t, err = time.Parse(time.DateOnly, startTime)
	t, err = time.Parse("2006-01-02 15:04", startTime)
	if err != nil {
		return
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
	return
}

func GetEndTime(endTime string) (t time.Time, err error) {
	//t, err = time.Parse(time.DateOnly, endTime)
	t, err = time.Parse("2006-01-02 15:04", endTime)
	if err != nil {
		return
	}
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, 0, time.Local)
	return
}
