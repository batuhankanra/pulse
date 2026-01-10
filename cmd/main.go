package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/batuhankanra/pulse.git/internal/config"
	httpclient "github.com/batuhankanra/pulse.git/internal/httpClient"

	"github.com/batuhankanra/pulse.git/internal/models"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	cmd := os.Args[1]
	switch cmd {
	case "url-add":
		if len(os.Args) < 4 {
			fmt.Println("usage: pulse url-add <key> <baseUrl>")
			return
		}
		key := os.Args[2]
		value := os.Args[3]
		cfg, _ := config.Load()
		cfg.URLs[key] = value
		config.Save(cfg)

		fmt.Println("added:", key, "->", value)
	case "url-list":
		cfg, _ := config.Load()
		if len(cfg.URLs) == 0 {
			fmt.Println("no urls defined")
			return
		}
		b, err := json.MarshalIndent(cfg.URLs, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(string(b))

	case "url-del":
		if len(os.Args) < 3 {
			fmt.Println("usage: pulse url-rm <key>")
			return
		}
		key := os.Args[2]

		cfg, _ := config.Load()
		delete(cfg.URLs, key)
		config.Save(cfg)
		fmt.Println("removed:", key)

	case "url-update":
		if len(os.Args) < 4 {
			fmt.Println("usage: pulse url-set <key> <newBaseUrl>")
			return
		}
		key := os.Args[2]
		val := os.Args[3]
		cfg, _ := config.Load()
		if _, ok := cfg.URLs[key]; !ok {
			fmt.Println("not found:", key)
			return
		}
		cfg.URLs[key] = val
		config.Save(cfg)
		fmt.Println("updated:", key, "->", val)

	case "req":
		if len(os.Args) < 4 {
			fmt.Println("usage: pulse req <method> <url> [headers...]")
			return
		}
		method := strings.ToUpper(os.Args[2])
		rawUrl := os.Args[3]
		rawHeader := os.Args[4:]
		cfg, _ := config.Load()
		url := resolve(rawUrl, cfg)
		headers := make(map[string]string)
		for _, h := range rawHeader {
			r := resolve(h, cfg)
			parts := strings.SplitN(r, ":", 2)
			if len(parts) != 2 {
				fmt.Println("invalid header:", r)
				return
			}
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])

		}
		err := httpclient.Do(method, url, headers, "")
		if err != nil {
			fmt.Println("error:", err)
			return
		}

	default:
		usage()
	}

}

func usage() {
	fmt.Println("usage:")
	fmt.Println("pulse url-add <key> <baseUrl>")
	fmt.Println("pulse url-list")
	fmt.Println("pulse url-rm <key>")
	fmt.Println("pulse url-set <key> <newBaseUrl>")
	fmt.Println("pulse req <method> <url> [headers...]")
	fmt.Println("examples:")
	fmt.Println("pulse url-add b https://httpbin.org")
	fmt.Println("pulse req get !b/get")

}

func resolve(arg string, cfg *models.Config) string {
	if !strings.HasPrefix(arg, "!") {
		return arg
	}
	key := strings.TrimPrefix(arg, "!")
	if strings.Contains(key, "/") {
		parts := strings.SplitN(key, "/", 2)
		if base, ok := cfg.URLs[parts[0]]; ok {
			return strings.TrimRight(base, "/") + "/" + parts[1]
		}
	}
	if base, ok := cfg.URLs[key]; ok {
		return base
	}
	if h, ok := cfg.Headers[key]; ok {
		return h
	}
	return arg
}
