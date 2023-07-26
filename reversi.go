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

func input() (int, int) {
	var n [2]int
	fmt.Scan(&n[0], &n[1])
	return n[0], n[1]
}

func initialize() *Board {
	tokens := make([][]int, 8)
	for i := 0; i < 8; i++ {
		tokens[i] = []int{0, 0, 0, 0, 0, 0, 0, 0}
	}

	// initialize board
	tokens[3][3] = 1
	tokens[4][4] = 1
	tokens[3][4] = 2
	tokens[4][3] = 2

	return &Board{
		tokens: tokens,
	}
}

func (b *Board) PrintBoard() {
	fmt.Println("  0 1 2 3 4 5 6 7")
	for i := 0; i < 8; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 8; j++ {
			switch b.tokens[i][j] {
			case 1:
				fmt.Print("● ")
			case 2:
				fmt.Print("○ ")
			default:
				fmt.Print("- ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) isMovable(x, y int) string {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return "reject"
	}
	if b.tokens[x][y] != 0 {
		return "reject"
	}
	return "path"
}

func (b *Board) getFlippedDiscs(row, col, player_now int) [][]int {
	directions := [][2]int{
		{0, 1},  // 右
		{0, -1}, // 左
		{1, 0},  // 下
		{-1, 0}, // 上
		{1, 1},  // 右下
		{-1, 1}, // 右上
		{1, -1}, // 左下		{-1, -1}, // 左上
	}

	flipped := [][]int{}

	for _, dir := range directions {
		flipped = append(flipped, b.isgetFlipped(row, col, dir[0], dir[1], player_now)...)
	}

	return flipped
}

func (b *Board) isgetFlipped(row, col, dx, dy, player_now int) [][]int {
	opp := 3 - player_now //今のplayer
	flipped := [][]int{}  //フリップしたい駒　格納
	x, y := row+dx, col+dy

	for x >= 0 && x < 8 && y >= 0 && y < 8 && b.tokens[x][y] == opp { //相手の駒を判定
		flipped = append(flipped, []int{x, y})
		x += dx
		y += dy
	}

	if x >= 0 && x < 8 && y >= 0 && y < 8 && b.tokens[x][y] == player_now { //自分の駒まで　判定
		return flipped
	}
	return [][]int{} //フリップできないなら空配列を返す
}

//func (b *Board) getfilpedDiscs(row, col int) [][] int {

// }
func hasValidMove(b *Board, player_now int) bool {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.tokens[i][j] == 0 && len(b.getFlippedDiscs(i, j, player_now)) > 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	board := initialize()
	player := 1
	for {
		board.PrintBoard()
		x, y := input()
		judge := board.isMovable(x, y)
		if judge == "reject" {
			fmt.Println("again")
			continue
		}
		// filpped := [][]int{}
		flipped := board.getFlippedDiscs(x, y, player)
		if len(flipped) == 0 {
			fmt.Println("again")
			continue
		}
		board.tokens[x][y] = player
		for _, disc := range flipped {
			board.tokens[disc[0]][disc[1]] = player
		}
		if !hasValidMove(board, player) {
			fmt.Println("game over！")
			break
		}
		board.PrintBoard()
		player = 3 - player
	}
}
