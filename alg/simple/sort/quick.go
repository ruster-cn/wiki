package sort

//quick 从小到大排序,
func quick(array []interface{}, compare Compare) {
	quickWork(array, compare, 0, len(array))
}

//DEBUG:`start 包含 end不包含`
func quickWork(array []interface{}, compare Compare, start, end int) {
	sentinel := start
	if end <= start {
		return
	}
	//index记录比sentinel小的数字此时的下标，再发现比sentinel小的就可以放到index后面
	index := start + 1
	for i := index; i < end; i++ {
		if compare(array[sentinel], array[i]) {
			swap(array, index, i)
			index++
		}
	}
	//此时index的是第一个比sentinel大的数字
	swap(array, sentinel, index-1)
	//start-> index-2 都是比sentianel小的数字，index-1 是sentinel
	quickWork(array, compare, start, index-1)
	quickWork(array, compare, index, end)
}
