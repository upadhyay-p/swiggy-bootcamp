package stringutil

import "strings"

func FullNameNakedReturn(f,l string) (full string, length int){
	full = strings.ToUpper(f+" "+l)
	length = len(full)
	return
}