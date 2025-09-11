package utils

// 定义状态位
const (
	InitialStatus        int8 = 0b0000 // 初始状态
	ImmediatePushSuccess int8 = 0b0001 // 即时推送成功状态，如果即时推送失败则还是初始状态，但并不会与第二天推送状态混淆，而是可以描述叠加状态
	SecondDayPushSuccess int8 = 0b0010 // 第二天推送成功状态
	SecondDayPushFail    int8 = 0b0100 // 第二天推送失败状态
)

// SetStatus 设置状态
func SetStatus(currentStatus int8, statusToSet int8) int8 {
	return currentStatus | statusToSet
}

// ClearStatus 清除状态
func ClearStatus(currentStatus int8, statusToClear int8) int8 {
	return currentStatus &^ statusToClear
}

// CheckStatus 检查状态
func CheckStatus(currentStatus int8, statusToCheck int8) bool {
	return currentStatus&statusToCheck != 0
}
