package main

import (
	"bufio"
	"fmt"
	"os"
	"pulse/internal/storage"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	urlStore, _ := storage.NewStorage()

	for {
		fmt.Print("pulse>")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}
		parts := strings.Split(line, " ")
		command := parts[0]
		switch command {
		case "url-add":
			if len(parts) < 3 {
				fmt.Println("Kullanım: url-add <key> <url>")
				continue
			}
			key := parts[1]
			value := parts[2]
			if err := urlStore.AddURL(key, value); err != nil {
				fmt.Printf("kaydedilmedi: %s", err)
				continue
			}
			fmt.Println("kaydedildi")
		case "url-list":
			if err := urlStore.ListURL(); err != nil {
				fmt.Println("error", err)
			}

		}

	}
}
