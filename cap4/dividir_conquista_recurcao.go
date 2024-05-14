package cap4

func soma(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + soma(arr[1:])
}
