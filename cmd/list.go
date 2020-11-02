package cmd

import (
	"fmt"
	"log"

	"github.com/qjvtenkroode/endless/pkg/endless"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all URLs from the Endless list",
	Long:  "List all URLs from the Endless list",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		store, err := endless.NewBoltEndlessStore()
		if err != nil {
			log.Fatalf("Main - failed to initiate a new store: %v", err)
		}
		e := endless.CreateEndless(store)
		items, err := e.List()
		if err != nil {
			log.Fatalf("Main - list failed: %v", err)
		}
		for _, i := range items {
			fmt.Printf("ID: %v\nURL: %v\nRead: %v\n\n", i.ID, i.Url, i.Read)
		}
	},
}
