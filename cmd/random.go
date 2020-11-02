package cmd

import (
	"fmt"
	"log"

	"github.com/qjvtenkroode/endless/pkg/endless"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(randomCmd)
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random URL",
	Long:  "Get a random URL from your Endless reading list",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := endless.NewBoltEndlessStore()
		if err != nil {
			log.Fatalf("Main - failed to initiate new store: %v", err)
		}
		e := endless.CreateEndless(store)
		item, err := e.Random()
		if err != nil {
			log.Fatalf("Main - random failed: %v", err)
		}
		fmt.Printf("ID: %v\nURL: %v\nRead: %v\n", item.ID, item.Url, item.Read)
	},
}
