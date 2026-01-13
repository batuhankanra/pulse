# Pulse CLI 🚀

Pulse, terminal üzerinden hızlı ve konfigüre edilebilir HTTP istekleri göndermenizi sağlayan modern bir CLI aracıdır.

## 🛠 Kullanım

### Komutlar

- **URL Yönetimi:**
  - `pulse url-add [key] [url]` - Yeni bir URL kısayolu ekler.
  - `pulse url-list` - Tüm kayıtlı URL'leri listeler.
  - `pulse url-del [key]` - Bir URL kısayolunu siler.

- **Header & Body Yönetimi:**
  - `pulse header-add [id] [k] [v]` - Header grubuna veri ekler.
  - `pulse header-list` - Kayıtlı header gruplarını listeler.
  - `pulse body-add [id] [k] [v]` - Body grubuna (JSON) veri ekler.
  - `pulse body-list` - Kayıtlı body gruplarını listeler.

- **İstek Gönderme (Request):**
  - `pulse req [method] [url] [opts]` - HTTP isteği gönderir.

### 💡 İpucu (Semboller)

- **`:` veya `@`** : Config'den veri çekmek için kullanılır.
- **Path Ekleme**: `:api/users` kullanımı, `api` değerinin sonuna `/users` ekler.
- **PowerShell Notu**: Windows PowerShell kullanıyorsanız `@` veya `:` içeren argümanları mutlaka tek tırnak içine alın: `'@a/1'`.

---

## 🚀 Örnekler

```bash
# 1. URL ve Auth Kaydet
pulse url-add api [https://jsonplaceholder.typicode.com](https://jsonplaceholder.typicode.com)
pulse header-add auth Authorization 'Bearer 123'
pulse body-add login user admin

# 2. GET İsteği (Path ile)
pulse req get ':api/users/1'

# 3. POST İsteği (Kısayollarla)
pulse req post ':api/posts' '@auth' '@login'

# 4. Manuel Header
pulse req get google.com User-Agent:Pulse