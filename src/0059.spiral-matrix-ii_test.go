package leetcode

import (
	"fmt"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	n := 4
	fmt.Printf("n = %d\n", n)
	result := generateMatrix(n)

	fmt.Println(result)
}
