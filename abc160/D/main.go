package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, X int, Y int) {
	m := countDist(N, X, Y)
	for i := 1; i < N; i++ {
		fmt.Println(m[i])
	}
}

func countDist(N, X, Y int) map[int]int {
	m := map[int]int{}
	for i := 1; i <= N; i++ {
		for j := i + 1; j <= N; j++ {
			dist := calcDist(i, j, X, Y)
			m[dist]++
		}
	}
	return m
}

func calcDist(i, j, X, Y int) int {
	normalDist := lib_AbsInt(i - j)
	XDist := lib_AbsInt(X - i)
	YDist := lib_AbsInt(Y - j)
	shortcutDist := XDist + YDist + 1
	XDist2 := lib_AbsInt(X - j)
	YDist2 := lib_AbsInt(Y - i)
	shortcutDist2 := XDist2 + YDist2 + 1
	return lib_MustMinInt([]int{normalDist, shortcutDist, shortcutDist2})
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int64
	scanner.Scan()
	N, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var X int64
	scanner.Scan()
	X, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	var Y int64
	scanner.Scan()
	Y, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	solve(int(N), int(X), int(Y))
}