package utils

// Intersect 求两个int类型slice的交集
func Intersect(slice1, slice2 []int) []int {
	intersectMap := make(map[int]bool)
	var intersection []int

	// 遍历第一个slice，记录每个元素
	for _, value := range slice1 {
		intersectMap[value] = true
	}

	// 遍历第二个slice，检查元素是否在map中
	for _, value := range slice2 {
		if _, exists := intersectMap[value]; exists {
			intersection = append(intersection, value)
			// 删除map中的元素，避免重复添加
			delete(intersectMap, value)
		}
	}

	return intersection
}

// Intersect 求两个int类型不同初
func Difference(slice1, slice2 []int) []int {
	diff := []int{}
	map2 := make(map[int]bool)
	for _, v := range slice2 {
		map2[v] = true
	}
	for _, v := range slice1 {
		if !map2[v] {
			diff = append(diff, v)
		}
	}
	return diff
}
