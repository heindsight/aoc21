package set

type Set interface {
	Add(interface{})
	Delete(interface{})
	Contains(interface{}) bool
	IsSubset(Set) bool
	Iter() chan interface{}
	Length() int
}

type set struct {
	members map[interface{}]struct{}
}

var exists = struct{}{}

func NewSet() *set {
	s := &set{}
	s.members = make(map[interface{}]struct{})
	return s
}

func (s *set) Add(value interface{}) {
	s.members[value] = exists
}

func (s *set) Delete(value interface{}) {
	delete(s.members, value)
}

func (s *set) Contains(value interface{}) bool {
	_, contains := (s.members)[value]
	return contains
}

func (s *set) IsSubset(other Set) bool {
	for value := range s.members {
		if !other.Contains(value) {
			return false
		}
	}
	return true
}

func (s *set) Iter() chan interface{} {
	out := make(chan interface{})

	go func() {
		for value := range s.members {
			out <- value
		}
		close(out)
	}()
	return out
}

func (s *set) Length() int {
	return len(s.members)
}
