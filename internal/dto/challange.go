package dto

type Challenge struct {
	Nonce      string `json:"nonce"`
	Difficulty int    `json:"difficulty"`
}
