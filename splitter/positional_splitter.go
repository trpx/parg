package splitter

type PositionalSplitter struct {
	opt Options
}

func (s *PositionalSplitter) Split(args []string) (parts Parts, remainder []string) {
	for idx, arg := range args {
		if s.opt.isPositionalArg(arg) {
			parts.Positional = append(parts.Positional, arg)
			remainder = args[idx+1:]
		} else {
			break
		}
	}
	return parts, remainder
}

func (s *PositionalSplitter) Next() Splitter {
	// todo
	return nil
}
