func hasDuplicate(nums []int) bool {
    mp := make(map[int]int)
	for _, val := range nums{
		mp[val] ++
	}
	for _,val := range mp{
		if val > 1 {
			return true
		}
	}
	return false
}
