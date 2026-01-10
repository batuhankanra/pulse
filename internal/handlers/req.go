package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/batuhankanra/pulse.git/internal/config"
	httpclient "github.com/batuhankanra/pulse.git/internal/httpClient"
	"github.com/batuhankanra/pulse.git/internal/models"
)

func resolveUrl(arg string, cfg *models.Config) string {
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
	return arg
}

func Req(args []string) {
	if len(args) < 3 {
		fmt.Println("usage: pulse req <method> <url> [options]")
		return
	}

	method := strings.ToUpper(args[1])
	rawURL := args[2]

	cfg, _ := config.Load()
	url := resolveUrl(rawURL, cfg)

	headers := make(map[string]string)
	var body string

	for i := 3; i < len(args); i++ {
		arg := args[i]

		// -header !c
		if arg == "-header" && i+1 < len(args) {
			key := strings.TrimPrefix(args[i+1], "!")
			i++

			if group, ok := cfg.Headers[key]; ok {
				for hk, hv := range group {
					headers[hk] = hv
				}
			} else {
				fmt.Println("header group not found:", key)
				return
			}
			continue
		}

		// -body !login
		if arg == "-body" && i+1 < len(args) {
			key := strings.TrimPrefix(args[i+1], "!")
			i++

			if group, ok := cfg.Body[key]; ok {
				b, err := json.Marshal(group)
				if err != nil {
					fmt.Println("body marshal error:", err)
					return
				}
				body = string(b)
			} else {
				fmt.Println("body group not found:", key)
				return
			}
			continue
		}

		// Serbest kullanım: !c veya !login
		if strings.HasPrefix(arg, "!") {
			key := strings.TrimPrefix(arg, "!")

			if group, ok := cfg.Headers[key]; ok {
				for hk, hv := range group {
					headers[hk] = hv
				}
				continue
			}

			if group, ok := cfg.Body[key]; ok {
				b, err := json.Marshal(group)
				if err != nil {
					fmt.Println("body marshal error:", err)
					return
				}
				body = string(b)
				continue
			}
		}

		// Normal "Key: Value" header
		r := resolveUrl(arg, cfg)
		parts := strings.SplitN(r, ":", 2)
		if len(parts) != 2 {
			fmt.Println("invalid header:", r)
			return
		}
		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	// JSON body varsa otomatik Content-Type
	if body != "" {
		if _, ok := headers["Content-Type"]; !ok {
			headers["Content-Type"] = "application/json"
		}
	}

	if err := httpclient.Do(method, url, headers, body); err != nil {
		fmt.Println("error:", err)
	}
}
