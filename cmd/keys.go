package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "List all keys",
	Run: func(cmd *cobra.Command, args []string) {
		for _, key := range RequestKeys() {
			fmt.Println(key)
		}
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}
