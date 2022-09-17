// Package kmp
// @Description: KMP算法，字符串匹配

package kmp

func firstAt(main, pattern []rune) int {
	// 模式串next数组
	patternNext := calNext(pattern)

	// 主串和模式串的长度
	mainLen, patternLen := len(main), len(pattern)

	// 主串和模式串的游标
	mainCursor, patternCursor := 0, 0

	for {
		// 2. 不匹配的情况下如何处理
		if pattern[patternCursor] != main[mainCursor] {
			if patternCursor == 0 {
				mainCursor++
			} else {
				// 数组的一个技巧，不需要加1，
				patternCursor = patternNext[patternCursor-1]
			}
		}

		// patternLen - (patternCursor + 1) ：模式串剩余匹配长度
		// mainCursor+1 ： 主串游标当前访问到的长度
		if (patternLen-(patternCursor+1))+(mainCursor+1) > mainLen {
			return -1
		}

		// 1. 比较模式串和主串比较
		if pattern[patternCursor] == main[mainCursor] {
			mainCursor++
			patternCursor++
		}

		// 全部匹配好了
		if patternCursor == patternLen {
			// 获取匹配到的第一位
			return mainCursor - patternCursor
		}
	}
}

//
// calNext
//  @Description: 计算next数组
//  @param pattern
//  @return []int
//
func calNext(pattern []rune) []int {
	next := make([]int, len(pattern))

	// todo

	return next
}
