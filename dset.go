package ds

// Simple union find data struction
// Utilizing an int splice to store
// the data
// Extendable and deletable, but usually the size is set at the start

type Dset struct {
	data []int
}

func MakeEmptyDSet()

func MakeDSet(num int) (r Dset) {
	r.data = make([]int, num)
	return
}

func (s Dset) Find(num int) int {
	return s.data[num]
}

func (s *Dset) Union(a, b int) {
	if s.Find(a) != s.Find(b) {
		replace := s.Find(a)
		with := s.Find(b)
		for val, _ := range s.data {
			if s.data[val] == replace {
				s.data[val] = with
			}
		}
	}
	return
}
