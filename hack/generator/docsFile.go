package generator

type DocsFile interface {
	Content() (string, error)
}
