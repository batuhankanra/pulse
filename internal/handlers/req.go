package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/batuhankanra/pulse.git/internal/config"
	httpclient "github.com/batuhankanra/pulse.git/internal/httpClient"
	"github.com/batuhankanra/pulse.git/internal/models"
)

func resolveUrl(raw string, cfg *models.Config) string {
	if strings.HasPrefix(raw, "@") {
		key := strings.TrimPrefix(raw, "@")
		if cfg.URLs != nil {
			if url, ok := cfg.URLs[key]; ok {
				return url
			}
		}
		fmt.Printf("warning: url alias '@%s' not found, using raw string\n", key)
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

	cfg, _ := config.Load() // Hata kontrolü eklenebilir

	// URL'i çözümle (@prod -> https://api...)
	url := resolveUrl(rawURL, cfg)

	headers := make(map[string]string)
	var body string

	// Prefix karakteri
	const prefix = "@"

	for i := 3; i < len(args); i++ {
		arg := args[i]

		// --- Durum 1: Spesifik Header Grubu (-header @common) ---
		if arg == "-header" && i+1 < len(args) {
			val := args[i+1]
			i++

			if strings.HasPrefix(val, prefix) {
				key := strings.TrimPrefix(val, prefix)
				if group, ok := cfg.Headers[key]; ok {
					// group burada map[string]string tipindedir
					for hk, hv := range group {
						headers[hk] = hv
					}
				} else {
					fmt.Println("header group not found:", key)
					return
				}
			} else {
				// Manuel giriş: -header "Auth: Bearer 123"
				parts := strings.SplitN(val, ":", 2)
				if len(parts) == 2 {
					headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
				}
			}
			continue
		}

		// --- Durum 2: Spesifik Body Grubu (-body @login) ---
		if arg == "-body" && i+1 < len(args) {
			val := args[i+1]
			i++

			if strings.HasPrefix(val, prefix) {
				key := strings.TrimPrefix(val, prefix)
				if group, ok := cfg.Body[key]; ok {
					// group map[string]string olduğu için JSON'a çeviriyoruz
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
				// Manuel body girişi: -body '{"foo":"bar"}'
				body = val
			}
			continue
		}

		// --- Durum 3: Kısayol Kullanımı (@auth veya @login gibi) ---
		if strings.HasPrefix(arg, prefix) {
			key := strings.TrimPrefix(arg, prefix)
			found := false

			// 1. Headers içinde ara ve ekle
			if cfg.Headers != nil {
				if group, ok := cfg.Headers[key]; ok {
					for hk, hv := range group {
						headers[hk] = hv
					}
					found = true
				}
			}

			// 2. Body içinde ara ve ekle
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

		// --- Durum 4: Standart Header (Key: Value) ---
		parts := strings.SplitN(arg, ":", 2)
		if len(parts) == 2 {
			headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
		}
	}

	// JSON body varsa Content-Type ekle
	if body != "" {
		if _, ok := headers["Content-Type"]; !ok {
			headers["Content-Type"] = "application/json"
		}
	}

	// İsteği yap
	if err := httpclient.Do(method, url, headers, body); err != nil {
		fmt.Println("error:", err)
	}
}
