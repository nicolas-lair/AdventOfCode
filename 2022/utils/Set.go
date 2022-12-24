package utils

type Set struct {
	content map[any]bool
}

func (s *Set) Has(item any) bool {
	_, ok := s.content[item]
	return ok
}

func (s *Set) Add(item any) bool {
	if s.content == nil {
		s.content = make(map[any]bool)
	}
	hasItem := s.Has(item)
	if !hasItem {
		s.content[item] = true
	}
	return hasItem
}

func (s *Set) Delete(item any) bool {
	hasItem := s.Has(item)
	if hasItem {
		delete(s.content, item)
	}
	return hasItem
}

func (s *Set) Intersection(other Set) *Set {
	newSet := new(Set)
	for k := range s.content {
		if other.Has(k) {
			newSet.Add(k)
		}
	}
	return newSet
}

func (s *Set) Union(other Set) *Set {
	newSet := new(Set)
	for k := range s.content {
		newSet.Add(k)
	}
	for k := range other.content {
		newSet.Add(k)
	}
	return newSet
}

func SetAsSlice[T any](s *Set) []T {
	var keys []T
	for k := range s.content {
		keys = append(keys, k.(T))
	}
	return keys
}

func PopulateSet[T any](items []T) Set {
	newSet := new(Set)
	for _, k := range items {
		newSet.Add(k)
	}
	return *newSet
}
