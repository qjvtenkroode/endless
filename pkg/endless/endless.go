package endless

import (
	"crypto/md5"
	"fmt"
)

type Item struct {
	ID   string
	Url  string
	Read bool
}

//type Endless struct {
//}

type Endless struct {
	store EndlessStore
}

type EndlessStore interface {
	Add(i *Item) error
	Get(s string) (*Item, error)
	List() ([]*Item, error)
	Random() (*Item, error)
}

func CreateEndless(store EndlessStore) *Endless {
	endlessList := store
	return &Endless{endlessList}
}

func CreateItem(url string) (*Item, error) {
	id := md5.Sum([]byte(url))
	i := Item{fmt.Sprintf("%x", id), url, false}
	return &i, nil
}

func (e Endless) Add(i *Item) error {
	err := e.store.Add(i)
	return err
}

func (e Endless) Get(s string) (*Item, error) {
	i, err := e.store.Get(s)
	return i, err
}

func (e Endless) List() ([]*Item, error) {
	i, err := e.store.List()
	return i, err
}

func (e Endless) Random() (*Item, error) {
	i, err := e.store.Random()
	return i, err
}
