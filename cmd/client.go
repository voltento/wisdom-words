package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func solvePoW(nonce string, difficulty int) string {
	prefix := strings.Repeat("0", difficulty)
	var counter int
	for {
		data := fmt.Sprintf("%s%d", nonce, counter)
		hash := sha256.Sum256([]byte(data))
		hashStr := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashStr, prefix) {
			return fmt.Sprintf("%d", counter)
		}
		counter++
	}
}

type challenge2 struct {
	Nonce      string `json:"nonce"`
	Difficulty int    `json:"difficulty"`
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)

	var challenge2 challenge2
	json.Unmarshal(buffer[:n], &challenge2)

	result := solvePoW(challenge2.Nonce, challenge2.Difficulty)
	conn.Write([]byte(result))

	n, _ = conn.Read(buffer)
	wisdom := string(buffer[:n])
	fmt.Println(wisdom)
}
