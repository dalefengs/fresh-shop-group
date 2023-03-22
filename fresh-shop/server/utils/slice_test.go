package utils

import (
	"fmt"
	"testing"
)

func TestSliceDifference(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	un := SliceDifference(slice1, slice2)
	fmt.Println("slice1 - slice2的差集为：", un)
}

func TestSliceUnionSet(t *testing.T) {
	slice1 := []string{"1", "2", "3", "6", "8"}
	slice2 := []string{"2", "3", "5", "0"}
	un := SliceUnionSet(slice1, slice2)
	fmt.Println("slice1，slice2的并集为：", un)
}

func TestSliceDifference1(t *testing.T) {
	db := []string{"1", "2", "3", "6", "8"}
	from := []string{"1", "2", "3", "6", "8"}
	un := SliceDifference(db, from)
	fmt.Println("删除数据：", un)
	un = SliceDifference(from, db)
	fmt.Println("添加数据：", un)
}
