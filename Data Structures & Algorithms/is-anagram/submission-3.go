import(
	"slices"
)

func sortStr(s string) (temp []rune){
	temp = []rune(s)
	slices.Sort(temp)
	return 
}
func isAnagram(s string, t string) bool {
	s1 := sortStr(s)
	s2 := sortStr(t)
	if len(s1) != len(s2){
		return false
	}
	for i := 0; i < len(s1); i ++{
		if s1[i] != s2[i]{
			return false
		}
	}
	return true
}
