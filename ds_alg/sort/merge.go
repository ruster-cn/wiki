package sort

func MergeSort(array []int, start, end int) []int {
	if end <= start {
		//退出
	}

	mid := (start + end) / 2

	pre := MergeSort(array, start, mid)
	last := MergeSort(array, mid, end)

	//合并结果

}

func merge(array []int, start, mid, end int) []int {
	return nil
}
