# 1. Go için temel imaj (Alpine ile optimize edilmiş)
FROM golang:1.23.4-alpine as build

# 2. Gerekli paketleri yükle (CGO ve SQLite desteği için)
RUN apk add --no-cache build-base gcc musl-dev git

# 3. Çalışma dizinini oluştur
WORKDIR /app

# 4. Go modüllerini ayrı kopyala (Cache kullanımı için)
COPY go.mod go.sum ./
RUN go mod tidy

# 5. Uygulama dosyalarını kopyala
COPY . .

# 6. CGO'yu etkinleştirerek Go uygulamasını derle
RUN CGO_ENABLED=1 go build -o main ./cmd/app

# 7. Küçük ve temiz bir imajda uygulamayı çalıştır
FROM alpine:latest

# 8. Redis bağlantısı için gerekli paketleri yükle
RUN apk --no-cache add ca-certificates

# 9. Uygulama dosyasını kopyala
COPY --from=build /app/main /app/main

# 10. Çalışma dizini
WORKDIR /app

# 11. Redis portu ve Go uygulaması portunu aç
EXPOSE 6379
EXPOSE 3000

# 12. Uygulamayı çalıştır
CMD ["./main"]
