package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int32
	arr = []int32{793810624, 895642170, 685903712, 623789054, 468592370}
	var arrTemp []int32
	arrTemp = arr
	//fmt.Println("Original Integers ==== :", arr)
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	arr = arr[1:]
	fmt.Println("Sorted Integers Asc == :", arr)
	var ascSum int64
	for _, number := range arr {
		ascSum += int64(number)
	}
	fmt.Printf("%d\n", ascSum)
	sort.Slice(arrTemp, func(i, j int) bool {
		return arrTemp[i] > arrTemp[j]
	})
	arrTemp = arrTemp[:len(arrTemp)-1]
	fmt.Println("Sorted Integers Desc = :", arrTemp)
	var descSum int64
	for _, number := range arr {
		descSum += int64(number)
	}
	fmt.Printf("%d\n", descSum)
}
