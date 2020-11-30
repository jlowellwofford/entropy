package main

import (
	"log"

	"github.com/jlowellwofford/entropy/pkg/entropy"
)

func main() {
	ent, err := entropy.GetEntropy()
	if err != nil {
		log.Fatalf("failed to get entropy: %v", err)
	}
	log.Printf("current entropy: %d", ent)
}
