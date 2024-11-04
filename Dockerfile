# Etapa de construcción
FROM golang:alpine AS builder
WORKDIR /app

# Instalar dependencias y copiar el código
COPY go.mod ./
COPY go.sum ./
RUN go mod download 
COPY . .

# Construir la aplicación
RUN go build -o /to-do-app 

# Etapa final (imagen ligera)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /to-do-app .
COPY ./.env .

# Descargar wait-for-it
RUN apk add --no-cache bash curl && \
    curl -o /wait-for-it.sh https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh && \
    chmod +x /wait-for-it.sh
    
EXPOSE 8080

# Ejecutar la aplicación
CMD ["./to-do-app"]
