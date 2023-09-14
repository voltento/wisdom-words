package main

import (
	"flag"
	"github.com/voltento/wisdom-words/internal/pow"
	"github.com/voltento/wisdom-words/internal/web/client"
	"github.com/voltento/wisdom-words/internal/wisdom"
)

func main() {
	networkFlag := flag.String("network", "tcp", "Network protocol to use (e.g., tcp, udp)")
	addressFlag := flag.String("address", "127.0.0.1:8080", "Address of the server to connect to")

	flag.Parse()
	println("Connecting to " + *addressFlag)

	c := wisdom.NewClient(client.NewTCPClient(*networkFlag, *addressFlag), pow.NewSHA())
	words := c.GetWisdom()

	println(words.Wisdom)
}
