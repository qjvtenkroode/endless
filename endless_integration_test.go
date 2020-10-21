package main

import (
	"reflect"
	"testing"
)

func TestEndlessIntegration(t *testing.T) {
	store, err := NewBoltEndlessStore()
	assertNoError(t, err)
	endless := CreateEndless(store)

	item, err := CreateItem("http://www.test.com")
	if err != nil {
		t.Errorf("Endless integration - failed to create item, got %v.", err)
	}

	err = endless.Add(item)
	if err != nil {
		t.Errorf("Endless integration - failed to add item, got %v.", err)
	}

	got, err := endless.Get(item.ID)
	if err != nil {
		t.Errorf("Endless integration - failed to get item, got %v.", err)
	}
	if !reflect.DeepEqual(got, item) {
		t.Errorf("Endless integration - got %v, want %v.", got, item)
	}

	items, err := endless.List()
	want := 1
	if err != nil {
		t.Errorf("Endless integration - failed to list items, got %v.", err)
	}
	if len(items) != want {
		t.Errorf("Endless integration - got %v, want %v.", len(items), want)
	}
}
