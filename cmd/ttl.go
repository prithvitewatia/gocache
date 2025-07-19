package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ttlCmd = &cobra.Command{
	Use:   "ttl <key>",
	Short: "Show remaining TTL of a key in seconds. -1 means no expiry",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: ttl <key>")
			return
		}
		key := args[0]
		ttl := RequestTtl(key)
		fmt.Printf("%d", ttl)
		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(ttlCmd)
}
