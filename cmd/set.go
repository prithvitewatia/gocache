package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

var setCmd = &cobra.Command{
	Use:   "set <key> <value> [ttl]",
	Short: "Set a key with value and optional TTL",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Usage: set <key> <value> [ttl]")
			return
		}
		key := args[0]
		value := args[1]
		ttl := time.Duration(0)
		if len(args) == 3 {
			parsedTTL, err := time.ParseDuration(args[2])
			if err != nil {
				fmt.Println("Invalid TTL format. Use formats like 5s, 1m, 2h")
				return
			}
			ttl = parsedTTL
		}
		cacheInstance.Set(key, value, ttl)
		fmt.Println("OK")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
