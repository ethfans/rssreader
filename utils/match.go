package utils

import (
	"rss-reader/globals"
	"strings"
)

/**
 * MatchStr函数用于检查输入字符串str是否满足特定条件：
 * 必须包含所有非减号开头的关键字（表示为'b类'关键词），同时不能包含任何以减号开头的关键字（表示为'a类'关键词）。
 * 关键字源自globals.MatchList，每个关键字可能包含多个由空格分隔的部分。
 * 当str满足上述条件时，会执行提供的回调函数，传递匹配到的原始字符串str给回调。
 *
 * @param str string 待检查的输入字符串。
 * @param callback func(string) 回调函数，当str满足条件时被调用，参数为匹配到的str。
 */
func MatchStr(str string, callback func(string)) {

	for _, v := range globals.MatchList {
		strFinal := strings.ToLower(strings.TrimSpace(str))
		v = strings.ToLower(strings.TrimSpace(v))
		parts := strings.Split(v, " ")
		hasB := false // 标记是否包含所有非减号开头的关键字
		hasA := false // 标记是否包含任何减号开头的关键字
		// 遍历关键字的各个部分进行检查
		for _, part := range parts {
			if strings.HasPrefix(part, "-") {
				// 检查是否有以减号开头的关键字在str中
				hasA = hasA || strings.Contains(strFinal, part[1:])
			} else {
				// 检查是否有非减号开头的关键字在str中
				hasB = hasB || strings.Contains(strFinal, part)
			}
		}

		// 如果包含所有'b类'关键字且不含任何'a类'关键字，则执行回调
		if hasB && !hasA {
			callback(str)
		}
	}
}
