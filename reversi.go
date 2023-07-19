package main

import "fmt"

const (
	Size    = 8
	player1 = 1
	player2 = 2
)

//make([][]int, Size)

type Board struct {
	tokens [][]int
}

// func (b *Board) put(x, y int, u string) {
// 	if b.tokens[x+4*y] != 0 {
// 		return
// 	} else if u == "o" {
// 		b.tokens[x+4*y] = 1
// 	} else {
// 		b.tokens[x+4*y] = -1
// 	}

// }
// func (b *Board) get(x, y int) string {
// 	if b.tokens[x+4*y] == 1 {
// 		return "o"
// 	}
// 	if b.tokens[x+4*y] == -1 {
// 		return "●"
// 	}
// 	return "other"
// }

func input() (int, int) {
	var n [2]int
	fmt.Scan(&n[0], &n[1])
	return n[0], n[1]
}

func initialize() *Board {
	tokens := make([][]int, 10)
	for i := 0; i < 10; i++ {
		tokens[i] = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	}

	// initialize board
	tokens[4][4] = 1
	tokens[5][5] = 1
	tokens[4][5] = -1
	tokens[5][4] = -1

	fmt.Println(tokens)
	return &Board{
		tokens: tokens,
	}

}

func main() {
	list := []int{0, 0, 0, 0, 0, 1, -1, 0, 0, -1, 1, 0, 0, 0, 0, 0}
	fmt.Println(list[0:4])
	fmt.Println(list[4:8])
	fmt.Println(list[8:12])
	fmt.Println(list[12:16])
}

// func (b *Board) isMovable(x, y int) error {
// 	if x <= 0 || x >= 9 || y <= 0 || y >= 9 {
// 		return fmt.Errorf("Invalid Move, out of range")
// 	}
// 	if Board[x][y] != 0 {
// 		return fmt.Errorf("Invalid Move, already placed")
// 	}

// }

// func (g *Game) getFlippedDir(row, col int) [][]int {
// 	directions := [][2]int{
// 		{0, 1},  // 右
// 		{0, -1}, // 左
// 		{1, 0},  // 下
// 		{-1, 0}, // 上
// 		{1, 1},  // 右下
// 		{-1, 1}, // 右上
// 		{1, -1}, // 左下		{-1, -1}, // 左上
// 	}

// 	flipped := [][]int{}

// 	for _, dir := range directions {
// 		flipped = append(flipped, b.isgetFlipped(row, col, dir[0], dir[1]))
// 	}

// 	return flipped
// }

// func (b *Board) isgetFlipped(row, col, dx, dy int) [][]int {
// 	opp := 3 - player_now //今のplayer
// 	flipped := [][]int{}  //フリップしたい駒　格納
// 	x, y := row+dx, col+dy

// 	for x >= 1 && x <= 8 && y >= 1 && y <= 8 && b.Board[x][y] == opp { //相手の駒を判定
// 		flipped = append(flipped, []int{x, y})
// 		x += dx
// 		y += dy
// 	}

// 	if x >= 1 && x <= 8 && y >= 1 && y <= 8 && b.Board[x][y] == player_now { //自分の駒まで　判定
// 		return flipped
// 	}
// 	return [][]int{} //フリップできないなら空配列を返す
// }
