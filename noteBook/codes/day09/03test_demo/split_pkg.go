package split_pkg

import (
	// "fmt"
	"strings"
)

func MySplit(s string, sep string) []string {
	var ret []string
	i := strings.Index(s, sep)
	for i > -1 {
		ret = append(ret, strings.TrimSpace(s[:i]))
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	ret = append(ret, s)

	return ret
}
