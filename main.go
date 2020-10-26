package main

import "github.com/qjvtenkroode/endless/cmd"

func main() {
	cmd.Execute()
	/*case "list":
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
	*/
}
