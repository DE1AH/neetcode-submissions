import(
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, val := range strs{
		temp := []rune(val)
		slices.Sort(temp)
		key := string(temp)
		mp[key] = append(mp[key], val)
	}
	res := make([][]string, 0)
	for _, val := range mp{
		res = append(res, val)
	}
	return res
}
