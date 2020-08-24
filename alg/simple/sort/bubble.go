package sort

//Compare 比较 a 和 b 大小
//	1. if a>b,return ture
//	2. if a<b,return false
type Compare func(a, b interface{}) bool

//bubble 从小到大排序
func bubble(array []interface{}, compare Compare) {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if compare(array[i], array[j]) {
				swap(array, i, j)
			}
		}
	}
}

func swap(array []interface{}, i, j int) {
	tmp := array[i]
	array[i] = array[j]
	array[j] = tmp
}
