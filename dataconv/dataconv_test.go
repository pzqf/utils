package dataconv

import (
	"fmt"
	"testing"
)

func TestConv(t *testing.T) {
	a, err := String2Int("10")
	fmt.Println("String2Int:", a, err)
}