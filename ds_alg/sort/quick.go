package sort

func QuickSort(array []int) []int {
	return quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, start, end int) []int {
	if start < end {
		index := partition(array, start, end)
		quickSort(array, start, index-1)
		quickSort(array, index+1, end)
	}
	return array
}

func partition(array []int, start, end int) int {
	//index 0 作为哨兵
	sentinal := array[start]

	//遍历后面的将大于 index 0的
	for start <= end {
		if array[start] < sentinal && array[end] >= sentinal {
			start++
			end--
			continue
		}
		swamp(array, start, end)
	}
	swamp(array, 0, start)
	return start
}

func swamp(array []int, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
