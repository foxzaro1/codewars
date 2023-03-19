package kata

func Grow(arr []int) int {
	multiply := 1
	for _, value := range arr {
		multiply *= value
	}
	return multiply
}
