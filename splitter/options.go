package splitter

type Options struct {
	prefix string

	longBoolArgNames  []string
	shortBoolArgNames []string

	longStringArgNames  []string
	shortStringArgNames []string

	subCommandNames []string
}

func (o *Options) isPositionalArg(arg string) bool {
	if arg[:len(o.prefix)] == o.prefix {
		return false
	}
	if o.isSubCommandArg(arg) {
		return false
	}
	return true
}

func (o *Options) isSubCommandArg(arg string) bool {
	for _, i := range o.subCommandNames {
		if i == arg {
			return true
		}
	}
	return false
}
