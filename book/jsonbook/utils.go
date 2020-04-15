package jsonbook

import "strings"

func Strings(strs ...string) string {
	return strings.Join(strs, "")
}
