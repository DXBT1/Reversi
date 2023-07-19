package main

import "fmt"

type Board struct {
	tokens [][]int
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

func input() (int, int) {
	var n [2]int
	fmt.Scan(&n[0], &n[1])
	return n[0], n[1]
}

func (b *Board) initial() string {
	for i < 10 {
		for j < 10 {
			if (i == 4 && j == 4) || (i == 5 && j == 5) {
				b.tokens[i][j] == 1
			} else if (i == 5 && j == 4) || (i == 4 && j == 5) {
				b.tokens[i][j] == -1
			} else if i == 0 || j == 0 || i == 9 || j == 9 {
				b.tokens[i][j] == 2
			} else {
				b.tokens[i][j] == 0
			}

		}
	}

}

func main() {
	list := []int{0, 0, 0, 0, 0, 1, -1, 0, 0, -1, 1, 0, 0, 0, 0, 0}
	fmt.Println(list[0:4])
	fmt.Println(list[4:8])
	fmt.Println(list[8:12])
	fmt.Println(list[12:16])
}
