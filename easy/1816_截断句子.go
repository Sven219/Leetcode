package main

import (
	"fmt"
)

// func truncateSentence(s string, k int) string {
// 	strarray := strings.Fields(strings.TrimSpace(s))
// 	l := len(strarray)
// 	strarray = strarray[0:k:l]
// 	str := strings.Join(strarray, ` `)
// 	return str
// }

func truncateSentence(s string, k int) string {
	n, end, count := len(s), 0, 0
	for i := 1; i <= n; i++ {
		if i == n || s[i] == ' ' {
			count++
			if count == k {
				end = i
				break
			}
		}
	}
	return s[:end]
}

func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Printf(truncateSentence(s, k))
}
