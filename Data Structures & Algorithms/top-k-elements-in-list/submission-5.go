func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	freq := make ([][]int, len(nums) + 1)
	for _, val := range nums{
		mp[val] ++
	}
	for key, cnt := range mp{
		freq[cnt] = append(freq[cnt], key)
	}
	res := make([]int, 0)
	for i := len(freq) - 1; i > 0; i --{
		for _,val := range freq[i]{
			res = append(res, val)
			if len(res) == k{
				return res
			}
		}
	}
	return res
}
