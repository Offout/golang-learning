package main

import (
	"sort"
	"strconv"
)

func main() {
	GetMapValuesSortedByKey(map[int]string{
		9:  "Сентябрь",
		1:  "Январь",
		2:  "Февраль",
		10: "Октябрь",
		5:  "Май",
		7:  "Июль",
		8:  "Август",
		12: "Декарь",
		3:  "Март",
		6:  "Июнь",
		4:  "Апрель",
		11: "Ноябрь",
	})
}

//Returns int
func ReturnInt() int {
	return 1
}

//Returns float
func ReturnFloat() float32 {
	return 1.1
}

//Returns int array
func ReturnIntArray() [3]int {
	return [...]int{1, 3, 4}
}

//Returns int slice
func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

//Makes int slice to string
func IntSliceToString(slice []int) string {
	var str string
	for _, val := range slice {
		str = str + strconv.FormatInt(int64(val), 10)
	}
	return str
}

//Merges slices
func MergeSlices(floatSlice []float32, intSlice []int32) []int {
	var tempSlice []int
	for _, val := range floatSlice {
		tempSlice = append(tempSlice, int(val))
	}
	for _, val := range intSlice {
		tempSlice = append(tempSlice, int(val))
	}
	return tempSlice
}

//Returns map values sorted by key
func GetMapValuesSortedByKey(sortableMap map[int]string) []string {
	var array []string
	var mapKeys []int

	for key := range sortableMap {
		mapKeys = append(mapKeys, int(key))
	}

	sort.Ints(mapKeys)

	for _, key := range mapKeys {
		array = append(array, sortableMap[key])
	}
	return array
}
