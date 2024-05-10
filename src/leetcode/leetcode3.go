package main

import (
	"fmt"
	"strings"
)

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10130 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {

	var tmp string
	res := 0

	for _, char := range s {
		if tmp != "" {
			index := strings.Index(tmp, string(char))

			if index == -1 {
				tmp += string(char)
			} else {
				if len(tmp) >= res {
					res = len(tmp)
				}
				tmp += string(char)
				tmp = tmp[index+1:]
			}
		} else {
			tmp += string(char)
		}
	}

	if len(tmp) >= res {
		res = len(tmp)
	}

	return res
}

func main() {
	s := " "
	length := lengthOfLongestSubstring(s)
	fmt.Println(length)
}

//leetcode submit region end(Prohibit modification and deletion)
