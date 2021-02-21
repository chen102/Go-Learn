//寻找最长不含有重复字符的子串
//算法：对于每个字母x
//		lastOccurred[x]不存在，或者<strt ------无需操作
//		lastOccurred[x] >=start -------更新start
//		更新lastOccurred[x]，更新maxLenght
//[解题思路]
//从左往右扫描，当遇到重复字母时，以上一个重复字母的index +1，作为新的搜索起始位置。这时需要清空Map，并且从index+1开始重新进行搜索。
package main

import "fmt"

func ll(s string) int {
	lastOccurred := make(map[byte]int) //记录每个字母最后出现的位置
	start := 0
	maxlengh := 0
	for i, ch := range []byte(s) { //[]byte():类型转换，每个字符都是byte

		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlengh {
			maxlengh = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlengh
}
func main() {
	fmt.Println(ll("bbbbb"))
	fmt.Println(ll("abcabcabcc"))
	fmt.Println(ll("wffeafacca"))
}
