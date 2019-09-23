package splitter

type Parts struct {
	Positional  []string            // {"arg-value1", "arg-value2", ...}
	NamedString map[string][]string // "flag-name": {"arg-value1", "arg-value2", ...}
	NamedBool   map[string][]bool   // same as the above but bool values
	SubCommand  []string            // args starting with subCommand name
	Tail        []string            // args after "--" arg
}

func (p *Parts) addNamedString(name string, value string) {
	_, ok := p.NamedString[name]
	if !ok {
		p.NamedString[name] = []string{}
	}
	p.NamedString[name] = append(p.NamedString[name], value)
}

func (p *Parts) addNamedBool(name string, value bool) {
	_, ok := p.NamedBool[name]
	if !ok {
		p.NamedBool[name] = []bool{}
	}
	p.NamedBool[name] = append(p.NamedBool[name], value)
}

func NewParts() (parts Parts) {
	parts.NamedString = make(map[string][]string)
	parts.NamedBool = make(map[string][]bool)
	return parts
}
