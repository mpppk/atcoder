package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(N int, A []int, Q int, X []int, Y []int) string {
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 4096
	const maxBufSize = 1000000
	scanner.Buffer(make([]byte, initialBufSize), maxBufSize)
	scanner.Split(bufio.ScanWords)
	var N int
	scanner.Scan()
	N, _ = strconv.Atoi(scanner.Text())
	A := make([]int, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		A[i], _ = strconv.Atoi(scanner.Text())
	}
	var Q int
	scanner.Scan()
	Q, _ = strconv.Atoi(scanner.Text())
	X := make([]int, Q)
	Y := make([]int, Q)
	for i := 0; i < Q; i++ {
		scanner.Scan()
		X[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		Y[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(solve(N, A, Q, X, Y))
}
