package goutils

func String2interfaceSlices(s []string) (i []interface{}) {
	i = make([]interface{}, len(s), len(s))
	for idx, _ := range s {
		i[idx] = s[idx]
	}
	return i
}
