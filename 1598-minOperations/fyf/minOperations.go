package fyf

func minOperations(logs []string) int {
	deep := 0
	for i := 0; i < len(logs); i++ {
		if logs[i] == "../" && deep == 0 {
			continue
		}
		if logs[i] == "../" {
			deep = deep - 1
			continue
		}
		if logs[i] == "./" {
			continue
		}
		deep++
	}
	return deep
}
