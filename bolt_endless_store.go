package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
)

func NewBoltEndlessStore() (*BoltEndlessStore, error) {
	db, err := bolt.Open("endless.db", 0600, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("readinglist"))
		if err != nil {
			return errors.New(fmt.Sprintf("Bolt - create bucket: %s", err))
		}
		return err
	})

	return &BoltEndlessStore{db}, err
}

type BoltEndlessStore struct {
	db *bolt.DB
}

func (s *BoltEndlessStore) Add(i *Item) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		var buff bytes.Buffer
		enc := gob.NewEncoder(&buff)
		err := enc.Encode(i)
		b := tx.Bucket([]byte("readinglist"))
		err = b.Put([]byte(i.ID), buff.Bytes())
		return err
	})
	return err
}

func (s *BoltEndlessStore) Get(str string) (*Item, error) {
	var item Item
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("readinglist"))
		v := b.Get([]byte(str))
		if v == nil {
			return errors.New("Bolt - item not stored")
		}
		buff := bytes.NewBuffer(v)
		dec := gob.NewDecoder(buff)
		err := dec.Decode(&item)
		return err
	})
	return &item, err
}

func (s *BoltEndlessStore) List() ([]*Item, error) {
	items := []*Item{}
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("readinglist"))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			var item Item
			buff := bytes.NewBuffer(v)
			dec := gob.NewDecoder(buff)
			err := dec.Decode(&item)
			if err != nil {
				return err
			}

			items = append(items, &item)
		}
		return nil
	})
	return items, err
}
