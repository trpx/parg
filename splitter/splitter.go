package splitter

type Splitter struct {
	stringArgNames  []string
	boolArgNames    []string
	subCommandNames []string
}

func (s *Splitter) Split(args []string) (parts Parts, remainder []string) {
	// todo
	return parts, remainder
}
