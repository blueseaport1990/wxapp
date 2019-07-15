// 其他工具包
package util

// 便利数组查找字符串索引位置，如果不存在则返回-1，使用完整遍历查找
func StringSearch(a []string, s string) int {
	for i, v := range a {
		if s == v {
			return i
		}
	}
	return -1
}

// 判断字符串是否在数组中
func StringInArray(a []string, s string) bool {
	return StringSearch(a, s) != -1
}

// 判断给定字符是否小写
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// 判断给定字符是否大写
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}
