/*
 * Copyright (c) 2021.  beihai@wingsxdu.com. All right reserved
 * ProjectName: $LeetCode
 * LastModified: 2021/1/24 下午5:03
 * Author: beihai
 */

package String

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射与电话按键相同。注意 1 不对应任何字母。
*/
var m = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	// 将第一个元素添加到切片中
	for _, v := range m[digits[0]] {
		res = append(res, string(v))
	}
	// 如果只有一个元素，则返回结果
	if len(digits) == 1 {
		return res
	}
	// 循环，依次添加
	for i := 1; i < len(digits); i++ {
		s := m[digits[i]]
		arr := res
		res = []string{}
		for _, v := range s {
			for i := 0; i < len(arr); i++ {
				res = append(res, arr[i]+string(v))
			}
		}
	}
	return res
}
