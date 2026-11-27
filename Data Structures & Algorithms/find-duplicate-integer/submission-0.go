func findDuplicate(nums []int) int {
    mp := make(map[int]int)
	for _, val := range nums{
		mp[val] ++
	}
	for key, val := range mp{
		if val > 1{
			return key
		}
	}
	return 0
}
