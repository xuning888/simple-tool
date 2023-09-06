package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {

	pNum, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	matrix := buildMatrix()
	wg := sync.WaitGroup{}
	wg.Add(pNum)
	now := time.Now()
	for i := 0; i < pNum; i++ {
		go func() {
			defer wg.Done()
			multiply(matrix, matrix)
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("cost : %vms \n", end.UnixMilli()-now.UnixMilli())
}

func multiply(A, B *[100][100]float64) {
	result := [100][100]float64{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result[i][j] = A[i][j] * B[i][j]
		}
	}
}

func buildMatrix() *[100][100]float64 {
	matrix := [100][100]float64{}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			matrix[i][j] = rand.Float64()
		}
	}
	return &matrix
}
