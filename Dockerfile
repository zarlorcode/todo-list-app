# Etapa de construcción
FROM golang:alpine AS builder
WORKDIR /app

# Instalar dependencias y copiar el código
COPY go.mod ./
COPY go.sum ./
RUN go mod download 
RUN go get gorm.io/driver/postgres 
RUN go get gorm.io/gorm
COPY . .

# Construir la aplicación
RUN go build -o /to-do-app 

# Etapa final (imagen ligera)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /to-do-app .
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./to-do-app"]
