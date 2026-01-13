package handlers

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"

	"github.com/batuhankanra/pulse.git/internal/config"
	httpclient "github.com/batuhankanra/pulse.git/internal/httpClient"
	"github.com/batuhankanra/pulse.git/internal/models"
)

const prefix = ":"

func resolveUrl(raw string, cfg *models.Config) string {
	if cleanStr, found := strings.CutPrefix(raw, prefix); found {
		parts := strings.SplitN(cleanStr, "/", 2)
		key := parts[0]
		if url, ok := cfg.URLs[key]; ok {
			if len(parts) > 1 {
				return strings.TrimSuffix(url, "/") + "/" + parts[1]
			}
			return url
		}
	}
	return raw

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

		if arg == "-header" && i+1 < len(args) {
			val := args[i+1]
			i++

			if strings.HasPrefix(val, prefix) {
				key := strings.TrimPrefix(val, prefix)
				if group, ok := cfg.Headers[key]; ok {
					maps.Copy(headers, group)
				} else {
					fmt.Println("header group not found:", key)
					return
				}
			} else {
				parts := strings.SplitN(val, ":", 2)
				if len(parts) == 2 {
					headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
				}
			}
			continue
		}

		if arg == "-body" && i+1 < len(args) {
			val := args[i+1]
			i++

			if key, found := strings.CutPrefix(val, prefix); found {
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
			} else {
				body = val
			}
			continue
		}
		if key, found := strings.CutPrefix(arg, prefix); found {
			found := false
			if cfg.Headers != nil {
				if group, ok := cfg.Headers[key]; ok {
					maps.Copy(headers, group)
					found = true
				}
			}
			if cfg.Body != nil {
				if group, ok := cfg.Body[key]; ok {
					b, err := json.Marshal(group)
					if err != nil {
						fmt.Println("body marshal error:", err)
						return
					}
					body = string(b)
					found = true
				}
			}

			if !found {
				fmt.Println("warning: config key not found in headers or body:", key)
			}
			continue
		}
		parts := strings.SplitN(arg, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}
	if body != "" {
		if _, ok := headers["Content-Type"]; !ok {
			headers["Content-Type"] = "application/json"
		}
	}
	if err := httpclient.Do(method, url, headers, body); err != nil {
		fmt.Println("error:", err)
	}
}
