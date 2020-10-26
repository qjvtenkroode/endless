package cmd

import (
	"log"

	"github.com/qjvtenkroode/endless/pkg/endless"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an URL to the Endless list",
	Long:  "Add an URL to the Endless list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		store, err := endless.NewBoltEndlessStore()
		if err != nil {
			log.Fatalf("Main - failed to initiate a new store: %v", err)
		}
		e := endless.CreateEndless(store)
		i, _ := endless.CreateItem(args[0])
		err = e.Add(i)
		if err != nil {
			log.Fatalf("Main - add failed: %v", err)
		}
	},
}
