package main

import (
	"fmt"
	"github.com/XANi/junkbox/mm-test/bot"
	"os"
)

func main() {
	client, err := bot.Connect(os.Getenv(`MM_URL`), os.Getenv(`MM_USER`), os.Getenv(`MM_PASS`))
	if err != nil {
		fmt.Printf("Failed connection: [%+v] [%+v]\n", err, client)
	} else {
		fmt.Println("OK")
	}
	_ = client
}
