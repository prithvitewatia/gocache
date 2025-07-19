// cmd/root.go
package cmd

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Conf struct {
	ServerHost string
	ServerPort string
}

func loadConf() *Conf {
	conf := &Conf{
		ServerHost: "localhost",
		ServerPort: "28109",
	}
	return conf
}

var rootCmd = &cobra.Command{
	Use:   "gocache",
	Short: "A minimal in-memory cache like Redis",
	Long:  `GoCache is a fast, simple in-memory key-value store with TTL support, written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(`
			╔════════════════════════════╗
			║        GOCACHE CLI         ║
			║   A blazing fast KV store  ║
			╚════════════════════════════╝
		`)
		fmt.Println("Welcome to gocache shell (type 'exit' to quit)")
		for {
			fmt.Print("gocache> ")
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}
			line = strings.TrimSpace(line)
			if line == "exit" || line == "quit" {
				fmt.Println("Bye!")
				return
			}
			args := strings.Fields(line)
			if len(args) == 0 {
				continue
			}
			switch args[0] {
			case "set":
				setCmd.Run(cmd, args[1:])
			case "get":
				getCmd.Run(cmd, args[1:])
			case "del":
				delCmd.Run(cmd, args[1:])
			case "keys":
				keysCmd.Run(cmd, args[1:])
			case "ttl":
				ttlCmd.Run(cmd, args[1:])
			case "flushall":
				flushAllCmd.Run(cmd, args[1:])
			default:
				fmt.Println("Unknown command:", args[0])
			}
		}
	},
}

var Client = &http.Client{}
var Config = loadConf()

func init() {
	rootCmd.PersistentFlags().StringVar(
		&Config.ServerHost, "serverhost", "localhost", "gocache server host")
	rootCmd.PersistentFlags().StringVar(
		&Config.ServerPort, "serverport", "28109", "gocache server port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
