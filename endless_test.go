package main

import (
	"errors"
	"reflect"
	"testing"
)

type StubEndlessStore struct {
	EndlessList map[string]*Item
}

func (s *StubEndlessStore) Add(i *Item) error {
	_, ok := s.EndlessList[i.ID]
	if ok {
		return errors.New("ID already exists")
	}
	s.EndlessList[i.ID] = i
	return nil
}

func (s *StubEndlessStore) Get(str string) (*Item, error) {
	_, ok := s.EndlessList[str]
	if !ok {
		return nil, errors.New("ID does not exist")
	}
	return s.EndlessList[str], nil
}

func (s *StubEndlessStore) List() ([]*Item, error) {
	items := []*Item{}
	for _, i := range s.EndlessList {
		items = append(items, i)
	}
	return items, nil
}

func TestEndless(t *testing.T) {

	t.Run("Endless object initialisation", func(t *testing.T) {
		got := CreateEndless(nil)
		want := &Endless{nil}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v.", got, want)
		}
	})

	t.Run("Endless CreateItem - one item", func(t *testing.T) {
		got, _ := CreateItem("http://www.test.com")
		want := &Item{"7330d2d5f820390054efbfb267b8639e", "http://www.test.com", false}

		assertItem(t, got, want)
	})

	t.Run("Endless CreateItem - two items", func(t *testing.T) {
		got1, _ := CreateItem("http://www.test.com")
		want1 := &Item{"7330d2d5f820390054efbfb267b8639e", "http://www.test.com", false}
		got2, _ := CreateItem("http://www.anothertest.com")
		want2 := &Item{"13276e25781a53ce373cba68f0637a42", "http://www.anothertest.com", false}

		assertItem(t, got1, want1)
		assertItem(t, got2, want2)
	})

}

func TestEndlessStore(t *testing.T) {
	e := CreateEndless(&StubEndlessStore{map[string]*Item{}})
	item, _ := CreateItem("http://www.test.com")

	t.Run("EndlessStore - add one item", func(t *testing.T) {
		err := e.Add(item)

		assertNoError(t, err)
	})

	t.Run("EndlessStore - add duplicate item", func(t *testing.T) {
		err := e.Add(item)
		want := errors.New("ID already exists")

		if !reflect.DeepEqual(err, want) {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("EndlessStore - get one item", func(t *testing.T) {
		got, err := e.Get("7330d2d5f820390054efbfb267b8639e")
		want, _ := CreateItem("http://www.test.com")

		assertNoError(t, err)
		assertItem(t, got, want)
	})

	t.Run("EndlessStore - list all items - total of one item", func(t *testing.T) {
		got, err := e.List()
		want := 1

		assertNoError(t, err)
		if len(got) != want {
			t.Errorf("got %v, want %v.", got, want)
		}
	})
}

func assertItem(t *testing.T, got *Item, want *Item) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v.", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got %v, want nil.", got)
	}
}
