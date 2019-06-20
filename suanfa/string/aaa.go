package stringcode

import "fmt"

func Test() {
	fmt.Println("hello world")
}

/**
最长子串
*/
func MaxZichuan(s string) int {

	len_s := len(s)

	// 1、检查字符串长度
	switch len_s {
	case 0:
		return 0
	case 1:
		return 1
	default:
		// map主要是用于查重
		m := make(map[string]int, 100)
		// 存放数据
		norm_s := make([]string, 0, 100)

		// 最长的字符串
		var max_s []string
		// 最长的长度
		max_l := 0

		for i := 0; i < len_s; i++ {
			t_s := string(s[i])

			//说明重复了
			if t := m[t_s]; t > 0 {
				if max_l < len(m) {
					max_l = len(m)
					max_s = norm_s
				}

				// 只要把 t_s 前面的剔除即可
				m, norm_s = del_head(m, norm_s, t_s)
			}

			norm_s = append(norm_s, t_s)
			m[t_s] = 1
		}

		if max_l < len(m) {
			max_l = len(m)
			max_s = norm_s
		}

		return len(max_s)
	}
}

//
func del_head(m map[string]int, strs []string, s string) (map[string]int, []string) {
	for k, v := range strs {
		delete(m, v)
		if v == s {
			return m, strs[k+1:]
		}
	}

	return m, strs
}

//==========================================最长公共前缀================================================
// https://leetcode-cn.com/problems/longest-common-prefix/comments/
func MaxQianzhui(strs []string) string {
	fmt.Println(strs)
	// 获取字符串数组长度
	strL := len(strs)

	switch strL {
	case 0:
		return ""
	case 1:
		return strs[0]
	default:
		return qianzhui(strs)
	}
}

func qianzhui(strs []string) string {
	len_s := len(strs[0])

	r := make([]byte, 0)

Xunhuan:
	for i := 0; i < len_s; i++ { //最长字符
		var t byte
		for _, str := range strs {
			if i >= len(str) {
				break Xunhuan
			}

			if t > 0 && t != str[i] {
				break Xunhuan
			}

			t = str[i]
		}

		r = append(r, byte(t))
	}

	return string(r)
}

// ================ 字符串的排列 ==========
// https://leetcode-cn.com/problems/permutation-in-string/

func CheckInclusion(s1 string, s2 string) bool {
	// 对字符串2变成map
	// 字符串1中字符对应再字符串2中的基a找到
	// 检查 基a 是否连续

	return true
}
