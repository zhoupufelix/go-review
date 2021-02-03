package merge

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T){
	nums := []int{1,5,8,9,7,5,6,1,1}
	orders := MergeSort(nums)
	fmt.Println(orders)
}
