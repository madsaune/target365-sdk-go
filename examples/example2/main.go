package main

import (
	"fmt"
	"log"
	"os"

	"github.com/madsaune/target365-sdk-go/environments"
	"github.com/madsaune/target365-sdk-go/services/outmessage"
)

func main() {
	token := os.Getenv("STREX_TOKEN")

	client := outmessage.NewClient(string(environments.BaseURLShared), token)

	msg, err := client.Get("<transactionID>")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Status:", msg.StatusCode)
}
