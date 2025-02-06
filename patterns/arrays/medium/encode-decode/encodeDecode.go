package arrays

type Transfer struct{}

func (t *Transfer) Encode(strs []string) string {
	return ""
}

func (t *Transfer) Decode(strs string) []string {
	return []string{""}
}
