import(
	"slices"
)

func strSort(s string) string{
	runes := []rune(s)
	slices.Sort(runes)
	return string(runes)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	s = strSort(s)
	t = strSort(t)
	
	for i := 0; i < len(s); i ++{
		if s[i] != t[i]{
			return false
		}
	}
	return true
}
