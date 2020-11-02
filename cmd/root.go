package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "endless",
	Short: "Endless stores the endless list of browser tabs you keep open for later reading",
	Long:  "Endless stores the endless list of browser tabs you keep open for later reading",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("testing cobra")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
