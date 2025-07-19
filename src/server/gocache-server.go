package main

import "github.com/prithvitewatia/gocache/src"

type Conf struct {
	Port string
}

func main() {
	cacheInstance := src.NewCache()
	conf := &Conf{Port: "28109"}

	server := src.NewServer(cacheInstance)
	server.Start(conf.Port)

}
