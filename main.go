package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	store, err := NewBoltEndlessStore()
	if err != nil {
		log.Fatalf("Main - failed to initiate a new store: %v", err)
	}
	e := CreateEndless(store)

	switch len(args) {
	case 0:
		fmt.Println("no subcommand given")
		return
	case 1:
		fmt.Println("no data for subcommand")
		return
	}

	switch args[0] {
	case "add":
		i, _ := CreateItem(args[1])
		err := e.Add(i)
		if err != nil {
			log.Fatalf("Main - add failed: %v", err)
		}
	case "get":
		item, err := e.Get(args[1])
		if err != nil {
			log.Fatalf("Main - get failed: %v", err)
		}
		fmt.Printf("ID: %v\nURL: %v\nRead: %v\n", item.ID, item.Url, item.Read)
	case "list":
		items, err := e.List()
		if err != nil {
			log.Fatalf("Main - list failed: %v", err)
		}
		for _, i := range items {
			fmt.Printf("ID: %v\nURL: %v\nRead: %v\n", i.ID, i.Url, i.Read)
		}
	default:
		fmt.Println("not a suitable choice")
	}
}
