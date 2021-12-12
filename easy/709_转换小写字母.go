package main

import (
	"fmt"
	"strings"
)

func toLowerCase(s string) string {
	lower := &strings.Builder{}
	lower.Grow(len(s))
	for _, ch := range s {
		if 65 <= ch && ch <= 90 {
			ch |= 32
		}
		lower.WriteRune(ch)
	}
	return lower.String()
}

// func toLowerCase(s string) string{
//     return strings.ToLower(s)
// }

// 大写变小写、小写变大写 : 字符 ^= 32;

// 大写变小写、小写变小写 : 字符 |= 32;

// 小写变大写、大写变大写 : 字符 &= -33

func main() {
	str := "How are you"
	fmt.Println(toLowerCase(str))

}
