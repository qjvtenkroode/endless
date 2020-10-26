package cmd

import (
	"fmt"
	"log"

	"github.com/qjvtenkroode/endless/pkg/endless"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an URL from the Endless list",
	Long:  "Get an URL from the Endless list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store, err := endless.NewBoltEndlessStore()
		if err != nil {
			log.Fatalf("Main - failed to initiate a new store: %v", err)
		}
		e := endless.CreateEndless(store)
		item, err := e.Get(args[0])
		if err != nil {
			log.Fatalf("Main - get failed: %v", err)
		}
		fmt.Printf("ID: %v\nURL: %v\nRead: %v\n", item.ID, item.Url, item.Read)
	},
}
