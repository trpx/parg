package splitter

type Splitter interface {
	Next() Splitter
	Split(args []string) (parts Parts, remainder []string)
}
