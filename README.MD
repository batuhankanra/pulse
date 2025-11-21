⚡ Pulse
Terminal Tabanlı API Test ve Geliştirme Ortamı

Pulse, geliştiriciler için Go ile yazılmış, hafif, hızlı ve tamamen terminal üzerinden çalışan bir HTTP istemcisi ve test aracıdır. Postman'in sunduğu özellikleri, terminalin hızı ve sadeliği ile birleştirir.

Grafik arayüzlerin (GUI) hantallığından kurtulun; API'lerinizin nabzını terminalden ölçün.

🚀 Özellikler
Hızlı İstekler: GET, POST, PUT, DELETE ve diğer HTTP metodları için anlık sorgular.

JSON Desteği: JSON verilerini otomatik formatlama (pretty-print) ve renklendirme.

Test Senaryoları: Yanıt süreleri, durum kodları (status codes) ve gövde (body) içeriği için yerleşik test mekanizmaları.

Koleksiyon Yönetimi: Sorgularınızı JSON dosyaları olarak kaydedin ve tekrar kullanın.

Scriptlenebilir: CI/CD süreçlerine entegre edilebilir yapı.

Kaynak Dostu: Electron tabanlı uygulamaların aksine sadece birkaç MB RAM kullanır.

📦 Kurulum
Projeyi klonlayın ve derleyin:

Bash

# Projeyi klonla
git clone https://github.com/KULLANICI_ADIN/pulse.git

# Klasöre git
cd pulse

# Uygulamayı derle
go build -o pulse main.go

# (Opsiyonel) Global kullanım için path'e taşı
sudo mv pulse /usr/local/bin/
🛠 Kullanım
Pulse, sezgisel bir komut satırı arayüzüne sahiptir.

1. Basit İstek Atma (GET)
Bir URL'e hızlıca istek atıp yanıtı görün:

Bash

pulse get https://jsonplaceholder.typicode.com/users/1
2. Veri Gönderme (POST)
JSON verisi ile giriş yapma isteği:

Bash

pulse post https://api.ornek.com/login \
  --body '{"email": "test@test.com", "password": "123"}' \
  --header "Content-Type: application/json"
3. Test Çalıştırma
Pulse, yanıtın beklentilerinizi karşılayıp karşılamadığını test edebilir.

Örneğin, durum kodunun 200 olmasını ve yanıt süresinin 500ms altında olmasını bekliyorsanız:

Bash

pulse test https://api.ornek.com/status --expect-status 200 --max-time 500ms
4. Kayıtlı Koleksiyonu Çalıştırma
koleksiyon.json dosyasındaki tüm istekleri sırayla çalıştırır:

Bash

pulse run koleksiyon.json
📂 Proje Yapısı
pulse/
├── cmd/            # Komut satırı argümanlarını işleyen kodlar (CLI logic)
├── internal/       # API istemcisi ve yardımcı fonksiyonlar
│   ├── client/     # HTTP isteklerini yapan motor
│   ├── tester/     # Test ve assertion mantığı
│   └── utils/      # JSON formatlayıcı ve renk kodları
├── collections/    # Örnek istek koleksiyonları
├── main.go         # Uygulamanın giriş noktası
└── README.md       # Dokümantasyon