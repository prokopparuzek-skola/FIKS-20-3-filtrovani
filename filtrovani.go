package main

import (
	"fmt"
	"os"
	"regexp"
)

const ALPHABET_LEN = 128

type letter struct {
	end   bool
	next  [ALPHABET_LEN]*letter
	edges []byte
}

func (let *letter) add(word string) {
	if len(word) == 0 {
		let.end = true
		return
	}
	c := word[0]
	if let.next[c] == nil {
		let.next[c] = new(letter)
		let.next[c].add(word[1:])
		if let.edges == nil {
			let.edges = make([]byte, 0)
		}
		let.edges = append(let.edges, c)
	} else {
		let.next[c].add(word[1:])
	}
}

func (let *letter) parse(regS *string) {
	switch len(let.edges) {
	case 0:
	case 1:
		if let.end {
			*regS += "("
			*regS += string(let.edges[0])
			let.next[let.edges[0]].parse(regS)
			*regS += ")?"
		} else {
			*regS += string(let.edges[0])
			let.next[let.edges[0]].parse(regS)
		}
	default:
		*regS += "("
		for i, e := range let.edges {
			*regS += "("
			*regS += string(e)
			let.next[e].parse(regS)
			*regS += ")"
			if i < len(let.edges)-1 {
				*regS += "|"
			}
		}
		*regS += ")"
		if let.end {
			*regS += "?"
		}
	}
}

func main() {
	var M, N, P int
	var table map[byte]byte = make(map[byte]byte)
	var tmp string
	var words *letter
	var regS string

	fmt.Scan(&M)
	for v := 0; v < M; v++ {
		fmt.Scan(&N, &P)
		words = new(letter)
		for i := 0; i < 10; i++ {
			var count int
			fmt.Scan(&count)
			fmt.Scan(&tmp)
			for j := 0; j < count; j++ {
				table[tmp[j]] = byte(i + 48)
			}
		}
		for i := 0; i < N; i++ {
			var count int
			var tmpB []byte
			fmt.Scan(&count)
			tmpB = make([]byte, count)
			fmt.Scan(&tmp)
			for j := 0; j < count; j++ {
				tmpB[j] = table[tmp[j]]
			}
			words.add(string(tmpB))
		}
		words.parse(&regS)
		fmt.Fprintln(os.Stderr, regS)
		reg := regexp.MustCompile(regS)
		for i := 0; i < P; i++ {
			var hum string
			var count int
			fmt.Scan(&count)
			fmt.Scanln(&hum)
			indexes := reg.FindIndex([]byte(hum))
			if indexes == nil {
				fmt.Println(-1)
			} else {
				fmt.Println(indexes[0])
			}
		}
	}
}
