package module1

func SumList(list []int) (result int) {
	for _, val := range list {
		result += val
	}
	return result
}

// Makes more sense and more efficient with large list sizes
func SumListWithRecursive(list []int) int {

	if len(list) == 0 {
		return 0
	}
	return list[0] + SumListWithRecursive(list[1:])
}
