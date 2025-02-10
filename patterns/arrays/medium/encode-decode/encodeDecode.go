package arrays

import (
	"strconv"
	"strings"
)

type Transfer struct{}

func (t *Transfer) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var sizes []string
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	return strings.Join(sizes, ",") + "#" + strings.Join(strs, "")
}

func (t *Transfer) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	parts := strings.SplitN(encoded, "#", 2)
	sizes := strings.Split(parts[0], ",")
	var res []string
	i := 0
	for _, sz := range sizes {
		if sz == "" {
			continue
		}
		length, _ := strconv.Atoi(sz)
		res = append(res, parts[1][i:i+length])
		i += length
	}
	return res
}
