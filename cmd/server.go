package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

func generateNonce() string {
	const letters = "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	for i := 0; i < 16; i++ {
		sb.WriteByte(letters[rand.Intn(len(letters))])
	}
	return sb.String()
}

func proofOfWork(nonce string, difficulty int) string {
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

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	fmt.Println("Server listening on 0.0.0.0:8080")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Connection established")

	nonce := generateNonce()
	challenge := challenge{
		Nonce:      nonce,
		Difficulty: 4,
	}

	challengeJson, _ := json.Marshal(challenge)
	conn.Write(challengeJson)

	buffer := make([]byte, 1024)
	n, _ := conn.Read(buffer)
	clientResponse := string(buffer[:n])

	if clientResponse == proofOfWork(nonce, challenge.Difficulty) {
		conn.Write([]byte("Here is your word of wisdom: Always be yourself."))
	} else {
		conn.Write([]byte("Failed PoW challenge."))
	}
}

type challenge struct {
	Nonce      string `json:"nonce"`
	Difficulty int    `json:"difficulty"`
}
