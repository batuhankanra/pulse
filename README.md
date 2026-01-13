# Pulse

**Pulse**, Go ile yazılmış, CLI tabanlı, hızlı ve hafif bir HTTP istemcisidir.  
Postman benzeri bir deneyimi terminalde sunmayı hedefler.

API testleri, debug, otomasyon ve script’ler içinde kolayca kullanılabilecek şekilde tasarlanmıştır.

## Özellikler

- GET, POST, PUT, DELETE vb. tüm HTTP metodlarını destekler
- Header ekleme
- JSON / raw body gönderme
- Response bilgilerini okunabilir formatta gösterme
- JSON response’ları pretty-print etme
- Ortam değişkenleri desteği (yakında)
- Request’leri dosya olarak saklama (yakında)




## Kullanım 

- Kullanım:
   pulse <komut> [argümanlar]

- URL Yönetimi:
   url-add <key> <url>          Yeni bir URL kısayolu ekler
   url-list                     Tüm kayıtlı URL'leri listeler
   url-del <key>                Bir URL kısayolunu siler

- Header & Body Yönetimi:
   header-add <id> <k> <v>      Header grubuna veri ekler
   header-list                  Kayıtlı header gruplarını listeler
   body-add <id> <k> <v>        Body grubuna (JSON) veri ekler
   body-list                    Kayıtlı body gruplarını listeler

- İstek Gönderme (Request):
   req <method> <url> [opts]    HTTP isteği gönderir

- İpucu (Semboller):
   : veya @                     Config'den veri çekmek için kullanılır
   ' ' (Tek tırnak)             PowerShell'de @ kullanırken tırnak içine alın

- Örnekler:
   pulse url-add api https://api.example.com
   pulse header-add auth Authorization 'Bearer 123'
   pulse body-add login user admin
   pulse req get :api/users             (URL sonuna path ekleme)
   pulse req post :api/login :auth :login
   pulse req get google.com User-Agent:Pulse


