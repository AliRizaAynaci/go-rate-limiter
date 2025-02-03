# 1. Temel Go imajını kullan
FROM golang:1.23.4-alpine as build

# 2. Çalışma dizinini oluştur
WORKDIR /app

# 3. Go modüllerini yükle
COPY go.mod go.sum ./
RUN go mod tidy

# 4. Uygulama dosyalarını kopyala
COPY . .

# 5. Go uygulamasını derle
RUN go build -o main ./cmd/app

# 6. Çalıştırılabilir dosyayı yeni bir imajda başlat
FROM alpine:latest

# 7. Redis bağlantısı için gerekli komutları yükle
RUN apk --no-cache add ca-certificates

# 8. Uygulama dosyasını kopyala
COPY --from=build /app/main /app/main

# 9. Çalışma dizini
WORKDIR /app

# 10. Redis portu ve Go uygulaması portunu aç
EXPOSE 6379
EXPOSE 3000

# 11. Uygulamayı çalıştır
CMD ["./main"]
