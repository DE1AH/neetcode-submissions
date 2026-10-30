func hasDuplicate(nums []int) bool {
    mp := make(map[int]int)
	for _, val := range nums{
		mp[val] ++
	}
	for _, v := range mp{
		if v > 1{
			return true
		}
	}
	return false
}
