package sort

func MergeSort(array []int) []int {
	return mergeSort(array)
}

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	//左闭右开集合
	pre := mergeSort(array[:mid])
	last := mergeSort(array[mid:])

	//合并结果
	return merge(pre, last)
}

func merge(pre []int, last []int) []int {
	var result []int
	var i, j = 0, 0
	for i < len(pre) && j < len(last) {
		if pre[i] <= last[j] {
			result = append(result, pre[i])
			i++
		} else {
			result = append(result, last[j])
			j++
		}
	}
	result = append(result, pre[i:]...)
	result = append(result, last[j:]...)
	return result
}
