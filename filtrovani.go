package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var M, N, P int
	var table map[byte]int = make(map[byte]int)
	var tmp string
	var words []string

	fmt.Scan(&M)
	for v := 0; v < M; v++ {
		fmt.Scan(&N, &P)
		words = make([]string, N)
		for i := 0; i < 10; i++ {
			var count int
			fmt.Scan(&count)
			fmt.Scan(&tmp)
			for j := 0; j < count; j++ {
				table[tmp[j]] = i
			}
		}
		for i := 0; i < N; i++ {
			var count int
			fmt.Scan(&count)
			fmt.Scan(&tmp)
			for j := 0; j < count; j++ {
				words[i] += fmt.Sprint(table[tmp[j]])
			}
			words[i] = "(" + words[i] + ")"
		}
		for i := 0; i < P; i++ {
			var hum string
			var count int
			fmt.Scan(&count)
			fmt.Scanln(&hum)
			reg := regexp.MustCompile(strings.Join(words, "|"))
			fmt.Println(reg.FindIndex([]byte(hum)))
		}

	}
}
