package _9word_search

// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
//
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring.
// The same letter cell may not be used more than once.
//
//
//
// Constraints:
//
//    m == board.length
//    n = board[i].length
//    1 <= m, n <= 6
//    1 <= word.length <= 15
//    board and word consists of only lowercase and uppercase English letters.
//
//
//
// Follow up: Could you use search pruning to make your solution faster with a larger board?

func exist(board [][]byte, word string) bool {
	return step01(board, word)
}

// 最直观的递归的寻找方式，有一点需要注意就是方格不能多次使用
//
// [结果]
// 性能一般
//
// [参考]
// - 各种判断来避免不可能的情况：长度、不包含的字符
// - 不需要借助额外的空间，使用特殊的字符直接在 board 中进行修改
func step01(board [][]byte, word string) bool {
	used := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]bool, len(board[0]))
	}

	x, y := len(board), len(board[0])
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if findAround(board, used, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

func findAround(board [][]byte, used [][]bool, word string, pos, x, y int) bool {
	if board[x][y] != word[pos] || used[x][y] {
		return false
	}
	if pos == len(word)-1 {
		return true
	}

	used[x][y] = true
	pos++
	if x > 0 && findAround(board, used, word, pos, x-1, y) {
		return true
	}
	if x < len(board)-1 && findAround(board, used, word, pos, x+1, y) {
		return true
	}
	if y > 0 && findAround(board, used, word, pos, x, y-1) {
		return true
	}
	if y < len(board[0])-1 && findAround(board, used, word, pos, x, y+1) {
		return true
	}
	used[x][y] = false
	return false
}
