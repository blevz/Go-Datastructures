package ds

type set struct {
	data map[string]bool
}

func EmptySet() (r set) {
	r.data = make(map[string]bool)
	return
}

func (s set) size() int {
	return len(s.data)
}

func (s *set) addElem(val string) {
	_, present := s.data[val]
	if !present {
		s.data[val] = true
	}
}

func (s *set) delElem(val string) {
	if s.Exist(val) {
		delete(s.data, val)
	}
}

func (s set) Exist(val string) bool {
	_, present := s.data[val]
	return present
}

func (s set) union(o set) (r set) {
	r = EmptySet()
	for v, _ := range s.data {
		r.addElem(v)
	}
	for v, _ := range o.data {
		r.addElem(v)
	}
	return
}

func (s set) intersect(o set) (r set) {
	r = EmptySet()
	for v, _ := range s.data {
		if o.Exist(v) {
			r.addElem(v)
		}
	}

	return
}
func (s set) difference(o set) (r set) {
	r = EmptySet()
	for v, _ := range s.data {
		if !o.Exist(v) {
			r.addElem(v)
		}
	}
	return
}
