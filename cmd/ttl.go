package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl <key>",
	Short: "Show remaining TTL of a key",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: ttl <key>")
			return
		}
		key := args[0]
		if ttl, ok := cacheInstance.TTL(key); ok {
			fmt.Println(ttl)
		} else {
			fmt.Println("(nil)")
		}
	},
}

func init() {
	rootCmd.AddCommand(ttlCmd)
}
