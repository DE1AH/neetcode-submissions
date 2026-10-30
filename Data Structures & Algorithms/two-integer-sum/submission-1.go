func twoSum(nums []int, target int) []int {
    mp := make(map[int]int)
	for ind, val := range nums{
		tmp := target - nums[ind]
		if val, ok := mp[tmp]; ok{
			return []int{val, ind}
		}
		mp[val] = ind
	}
	return nil
}
