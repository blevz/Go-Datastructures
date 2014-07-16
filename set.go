package ds

//Data structure to represent a mathematical Set
type Set struct {
	data map[interface{}]bool
}

//Public Fucntions

func MakeEmptySet() *Set {
	var r Set
	r.data = make(map[interface{}]bool)
	return &r
}

func (s Set) Empty() bool {
	return len(s.data) == 0
}

func (s Set) Size() int {
	return len(s.data)
}

func (s *Set) AddElem(val interface{}) {
	_, present := s.data[val]
	if !present {
		s.data[val] = true
	}
}

func (s *Set) DelElem(val interface{}) {
	if s.Exist(val) {
		delete(s.data, val)
	}
}

func (s Set) Exist(val interface{}) bool {
	_, present := s.data[val]
	return present
}

func (s Set) Union(o *Set) *Set {
	r := MakeEmptySet()
	for v, _ := range s.data {
		r.AddElem(v)
	}
	for v, _ := range o.data {
		r.AddElem(v)
	}
	return r
}

func (s Set) Intersect(o *Set) *Set {
	r := MakeEmptySet()
	for v, _ := range s.data {
		if o.Exist(v) {
			r.AddElem(v)
		}
	}
	return r
}

func (s Set) Difference(o *Set) *Set {
	r := MakeEmptySet()
	for v, _ := range s.data {
		if !o.Exist(v) {
			r.AddElem(v)
		}
	}
	return r
}
