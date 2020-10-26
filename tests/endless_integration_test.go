package tests

import (
	"reflect"
	"testing"

	"github.com/qjvtenkroode/endless/pkg/endless"
)

func TestEndlessIntegration(t *testing.T) {
	store, err := endless.NewBoltEndlessStore()
	assertNoError(t, err)
	e := endless.CreateEndless(store)

	item, err := endless.CreateItem("http://www.test.com")
	if err != nil {
		t.Errorf("Endless integration - failed to create item, got %v.", err)
	}

	err = e.Add(item)
	if err != nil {
		t.Errorf("Endless integration - failed to add item, got %v.", err)
	}

	got, err := e.Get(item.ID)
	if err != nil {
		t.Errorf("Endless integration - failed to get item, got %v.", err)
	}
	if !reflect.DeepEqual(got, item) {
		t.Errorf("Endless integration - got %v, want %v.", got, item)
	}

	items, err := e.List()
	want := 1
	if err != nil {
		t.Errorf("Endless integration - failed to list items, got %v.", err)
	}
	if len(items) != want {
		t.Errorf("Endless integration - got %v, want %v.", len(items), want)
	}
}
