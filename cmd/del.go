package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del <key>",
	Short: "Delete a key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: del <key>")
			return
		}
		key := args[0]
		cacheInstance.Delete(key)
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
}
