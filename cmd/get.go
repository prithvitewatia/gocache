package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get the value of a key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: get <key>")
			return
		}
		key := args[0]
		if val, ok := RequestGet(key); ok {
			fmt.Println(val)
		} else {
			fmt.Println("(nil)")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
