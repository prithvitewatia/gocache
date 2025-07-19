package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var flushAllCmd = &cobra.Command{
	Use:   "flushall",
	Short: "Clear all keys from the cache",
	Run: func(cmd *cobra.Command, args []string) {
		RequestFlushAll()
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(flushAllCmd)
}
