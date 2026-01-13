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

	case "req":
		handlers.Req(os.Args[1:])

	default:
		usage()
	}

}

func usage() {
	fmt.Println("Pulse CLI - Hızlı HTTP İstemcisi")
	fmt.Println("\nKullanım:")
	fmt.Println("  pulse <komut> [argümanlar]")

	fmt.Println("\nURL Yönetimi:")
	fmt.Println("  url-add <key> <url>          Yeni bir URL kısayolu ekler")
	fmt.Println("  url-list                     Tüm kayıtlı URL'leri listeler")
	fmt.Println("  url-del <key>                Bir URL kısayolunu siler")

	fmt.Println("\nHeader & Body Yönetimi:")
	fmt.Println("  header-add <id> <k> <v>      Header grubuna veri ekler")
	fmt.Println("  header-list                  Kayıtlı header gruplarını listeler")
	fmt.Println("  body-add <id> <k> <v>        Body grubuna (JSON) veri ekler")
	fmt.Println("  body-list                    Kayıtlı body gruplarını listeler")

	fmt.Println("\nİstek Gönderme (Request):")
	fmt.Println("  req <method> <url> [opts]    HTTP isteği gönderir")

	fmt.Println("\nİpucu (Semboller):")
	fmt.Println("  : veya @                     Config'den veri çekmek için kullanılır")
	fmt.Println("  ' ' (Tek tırnak)             PowerShell'de @ kullanırken tırnak içine alın")

	fmt.Println("\nÖrnekler:")
	fmt.Println("  pulse url-add api https://api.example.com")
	fmt.Println("  pulse header-add auth Authorization 'Bearer 123'")
	fmt.Println("  pulse body-add login user admin")
	fmt.Println("  pulse req get :api/users             (URL sonuna path ekleme)")
	fmt.Println("  pulse req post :api/login :auth :login")
	fmt.Println("  pulse req get google.com User-Agent:Pulse")
}
