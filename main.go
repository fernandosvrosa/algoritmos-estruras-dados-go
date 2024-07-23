package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 6)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}

	fmt.Println(removeElement2(nums, 2))

}

func removeElement2(nums []int, val int) int {

	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}

	return j
}

func removeElement(nums []int, val int) int {
	count := 0
	size := len(nums) - 1
	var temp []int
	for i := 0; i <= size; i++ {
		if nums[i] == val {
			count++
		} else {
			temp = append(temp, nums[i])
		}
	}
	nums = temp

	return count
}

func rotate(nums []int, k int) {
	n := len(nums)
	fmt.Println(k % n)
	k = k % n // In case k is greater than the length of nums

	if k == 0 {
		return // No need to rotate if k is 0
	}
	numb := n - k
	arr1 := nums[:numb]
	arr2 := nums[numb:]

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	// Create a new array to store the result
	result := append(arr2, arr1...)

	// Copy the result back into nums
	copy(nums, result)
}
