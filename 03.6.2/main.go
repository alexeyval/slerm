package main

func main() {
	nums := [...]int{1, 2, 4, 4, 2, 3, 2, 5, 3}

	funcName(nums)
}

func funcName(nums [9]int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
