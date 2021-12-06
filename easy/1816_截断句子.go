package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	strarray := strings.Fields(strings.TrimSpace(s))
	l := len(strarray)
	strarray = strarray[0:k:l]
	str := strings.Join(strarray, ` `)
	return str
}

func main() {
	s := "Hello how are you Contestant"
	k := 4
	fmt.Printf(truncateSentence(s, k))
}
