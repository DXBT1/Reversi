package main

import "fmt"

type Board struct {
	tokens []int
}

func (b *Board) put(x, y int, u string) {
	if b.tokens[x+4*y] != 0 {
		return
	} else if u == "o" {
		b.tokens[x+4*y] = 1
	} else {
		b.tokens[x+4*y] = -1
	}

}
func (b *Board) get(x, y int) string {
	if b.tokens[x+4*y] == 1 {
		return "o"
	}
	if b.tokens[x+4*y] == -1 {
		return "●"
	}
	return "other"
}

func main() {
	list := []int{0, 0, 0, 0, 0, 1, -1, 0, 0, -1, 1, 0, 0, 0, 0, 0}
	fmt.Println(list[0:4])
	fmt.Println(list[4:8])
	fmt.Println(list[8:12])
	fmt.Println(list[12:16])
}
