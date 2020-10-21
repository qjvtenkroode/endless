package main

import "testing"

func TestBoltEndlessStore(t *testing.T) {

	t.Run("Bolt - connect to a bolt db", func(t *testing.T) {
		store, err := NewBoltEndlessStore()
		defer store.db.Close()

		assertNoError(t, err)
	})

	t.Run("Bolt - add single item", func(t *testing.T) {
		store, err := NewBoltEndlessStore()
		defer store.db.Close()

		assertNoError(t, err)

		item := &Item{"7330d2d5f820390054efbfb267b8639e", "http://www.test.com", false}
		err = store.Add(item)
		assertNoError(t, err)
	})

	t.Run("Bolt - get single item", func(t *testing.T) {
		store, err := NewBoltEndlessStore()
		defer store.db.Close()

		assertNoError(t, err)
		item := &Item{"7330d2d5f820390054efbfb267b8639e", "http://www.test.com", false}
		got, err := store.Get("7330d2d5f820390054efbfb267b8639e")
		assertNoError(t, err)
		assertItem(t, got, item)
	})
}
