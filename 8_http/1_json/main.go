package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Request struct {
	Login  string  `json:"login"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

type Response struct {
	Ok      bool    `json:"ok"`
	Balance float64 `json:"balance"`
}

var data = `
{
  "user": "Scrooge McDuck",
  "type": "deposit",
  "amount": 123.4
}
`

func main() {
	// simulate socket
	reader := strings.NewReader(data)
	// decode request
	dec := json.NewDecoder(reader)

	var req Request
	if err := dec.Decode(&req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Printf("got: %+v\n", req)

	// create response
	prevBalance := 1_000_000.0 // Loaded from database
	resp := map[string]any{
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}

	// encode respose
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(resp); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}
