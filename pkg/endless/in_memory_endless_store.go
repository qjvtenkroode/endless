package endless

import "errors"

func NewInMemoryEndlessStore() *InMemoryEndlessStore {
	return &InMemoryEndlessStore{map[string]*Item{}}
}

type InMemoryEndlessStore struct {
	EndlessList map[string]*Item
}

func (s *InMemoryEndlessStore) Add(i *Item) error {
	_, ok := s.EndlessList[i.ID]
	if ok {
		return errors.New("ID already exists")
	}
	s.EndlessList[i.ID] = i
	return nil
}

func (s *InMemoryEndlessStore) Get(str string) (*Item, error) {
	_, ok := s.EndlessList[str]
	if !ok {
		return nil, errors.New("ID does not exist")
	}
	return s.EndlessList[str], nil
}
