package game

type void struct{}

type set struct {
	values map[string]void
}

func newSet(arr []string) set {
	s := set{make(map[string]void)}
	for _, x := range arr {
		s.add(x)
	}
	return s
}

func (s *set) add(val string) {
	s.values[val] = void{}
}

func (s *set) delete(val string) {
	delete(s.values, val)
}

func (s *set) has(val string) bool {
	for k := range s.values {
		if k == val {
			return true
		}
	}
	return false
}

func (s *set) size() int {
	return len(s.values)
}

func (s *set) elements() []string {
	var v []string
	for k := range s.values {
		v = append(v, k)
	}
	return v
}
