package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/batuhankanra/pulse.git/internal/config"
	"github.com/batuhankanra/pulse.git/internal/handlers"
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
	case "header-add":
		if len(os.Args) < 5 {
			fmt.Println("usage: pulse header-add <id> <Header> <Value>")
			return
		}
		id := os.Args[2]
		key := os.Args[3]
		val := os.Args[4]
		cfg, _ := config.Load()
		if cfg.Headers[id] == nil {
			cfg.Headers[id] = make(map[string]string)
		}

		cfg.Headers[id][key] = val
		config.Save(cfg)

		fmt.Println("added header:", id, key, "->", val)
	case "header-list":
		cfg, _ := config.Load()
		if len(cfg.Headers) == 0 {
			fmt.Println("{}")
			return
		}
		b, err := json.MarshalIndent(cfg.Headers, "", " ")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(string(b))
	case "header-del":
		if len(os.Args) < 3 {
			fmt.Println("usage: pulse header-del <key>")
			return
		}
		key := os.Args[2]
		cfg, _ := config.Load()
		delete(cfg.Headers, key)
		config.Save(cfg)
		fmt.Println("Removed header:", key)
	case "header-set":
		if len(os.Args) < 4 {
			fmt.Println("usage: pulse header-set <key> \"Header: Value\"")
			return
		}
		id := os.Args[2]
		key := os.Args[3]
		val := os.Args[4]
		cfg, _ := config.Load()
		if _, ok := cfg.Headers[key]; !ok {
			fmt.Println("not found:", key)
			return
		}
		cfg.Headers[id][key] = val
		config.Save(cfg)
		fmt.Println("updated header:", key, "->", val)
	case "body-add":
		if len(os.Args) < 5 {
			fmt.Println("usage: pulse body-add <id> <key> <value>")
			return
		}

		id := os.Args[2]
		key := os.Args[3]
		val := os.Args[4]

		cfg, _ := config.Load()

		if cfg.Body[id] == nil {
			cfg.Body[id] = make(map[string]string)
		}

		cfg.Body[id][key] = val
		config.Save(cfg)

		fmt.Println("added body field:", id, key, "->", val)

	case "body-list":
		cfg, _ := config.Load()
		if len(cfg.Body) == 0 {
			fmt.Println("{}")
			return
		}
		b, err := json.MarshalIndent(cfg.Body, "", "  ")
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		fmt.Println(string(b))

	case "body-del":
		if len(os.Args) < 3 {
			fmt.Println("usage: pulse body-del <key>")
			return
		}
		id := os.Args[2]
		cfg, _ := config.Load()
		delete(cfg.Body, id)
		config.Save(cfg)
		fmt.Println("Removed body:", id)

	case "body-set":
		if len(os.Args) < 4 {
			fmt.Println("usage: pulse body-set <key> \"body: Value\"")
			return
		}
		id := os.Args[2]
		key := os.Args[3]
		val := os.Args[4]
		cfg, _ := config.Load()
		if _, ok := cfg.Body[key]; !ok {
			fmt.Println("not found:", key)
			return
		}
		cfg.Headers[id][key] = val
		config.Save(cfg)
		fmt.Println("updated body:", key, "->", val)
	case "req":
		handlers.Req(os.Args[1:])

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
