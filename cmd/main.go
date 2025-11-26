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
	store, _ := storage.NewStorage()

	for {
		fmt.Print("pulse> ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		// Komut parçalama
		parts := strings.Fields(line)
		command := parts[0]

		switch command {

		// ---------------- URL ADD ----------------
		case "url-add":
			if len(parts) < 3 {
				fmt.Println("Kullanım: url-add <key> <url>")
				continue
			}
			key := parts[1]
			// Çok kelimeli URL'ler için geri kalanı birleştir
			value := strings.Join(parts[2:], " ")

			if err := store.AddURL(key, value); err != nil {
				fmt.Println("Kaydedilemedi:", err)
				continue
			}
			fmt.Println("✔ URL kaydedildi")

		// ---------------- URL LIST ----------------
		case "url-list":
			store.ListURL()

		// ---------------- URL REMOVE ----------------
		case "url-del":
			if len(parts) < 2 {
				fmt.Println("Kullanım: url-remove <key>")
				continue
			}
			if err := store.DeleteURL(parts[1]); err != nil {
				fmt.Println("Silinemedi:", err)
				continue
			}
			fmt.Println("✔ URL silindi")

		// ---------------- HEADER ADD ----------------
		case "header-add":
			if len(parts) < 4 {
				fmt.Println("Kullanım: header-add <setName> <key> <value>")
				continue
			}
			setName := parts[1]
			key := parts[2]
			value := strings.Join(parts[3:], " ")

			if err := store.AddHeader(setName, key, value); err != nil {
				fmt.Println("Kaydedilemedi:", err)
				continue
			}
			fmt.Println("✔ Header kaydedildi")

		// ---------------- HEADER LIST ----------------
		case "header-list":
			store.ListHeader()

		// ---------------- HEADER REMOVE ----------------
		case "header-del":
			if len(parts) < 3 {
				fmt.Println("Kullanım: header-remove <setName> <key>")
				continue
			}
			setName := parts[1]
			key := parts[2]

			if err := store.RemoveHeader(setName, key); err != nil {
				fmt.Println("Silinemedi:", err)
				continue
			}
			fmt.Println("✔ Header silindi")

		// ---------------- HELP ----------------
		case "help":
			fmt.Println("Komutlar:")
			fmt.Println("url-add <key> <url>")
			fmt.Println("url-list")
			fmt.Println("url-list")
			fmt.Println("url-remove <key>")
			fmt.Println("header-add <setName> <key> <value>")
			fmt.Println("header-list")
			fmt.Println("header-remove <setName> <key>")
			fmt.Println("exit")
		// ---------------- REQUEST ------------- req method url -header -body
		case "req":
			if len(parts) < 3 {
				fmt.Println("Kullanım: req <method> <urlOrKey>/path [-header <headerSet>] [-body <jsonBody>]")
				continue
			}
			method := strings.ToUpper(parts[1])
			rawURL := parts[2]
			var headerSet string
			var body string

			for i := 3; i < len(parts); i++ {
				switch parts[i] {
				case "-header":
					if i+1 < len(parts) {
						headerSet = parts[i+1]
					}
				case "-body":
					if i+1 < len(parts) {
						body = parts[i+1]
					}
				}
			}
			finalURL := rawURL
			if strings.HasPrefix(rawURL, "!") {
				parts2 := strings.SplitN(strings.TrimPrefix(rawURL, "!"), "/", 2)
				key := parts2[0]
				baseUrl, ok := store.URLs[key]
				if !ok {
					fmt.Println("Url bulunamadı")
					continue
				}
				if len(parts2) == 2 {
					finalURL = baseUrl + "/" + parts2[1]

				} else {
					finalURL = baseUrl
				}
			}
			headers := map[string]string{}
			if headerSet != "" {
				if set, ok := store.Headers[headerSet]; ok {
					headers = set
				} else {
					fmt.Println("header bulunamadı", headerSet)
				}
			}
			if err := storage.SendRequest(finalURL, method, headers, body); err != nil {
				fmt.Println("istek hatası:", err)
			}

		// ---------------- EXIT ----------------
		case "exit":
			fmt.Println("Çıkılıyor...")
			return

		// ---------------- UNKNOWN COMMAND ----------------
		default:
			fmt.Println("Bilinmeyen komut:", command)
			fmt.Println("Yardım için: help")
		}
	}
}
