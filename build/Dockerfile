# Etap budowania aplikacji Go
FROM golang:1.22 AS builder

WORKDIR /app

COPY .. .

RUN go mod tidy
RUN go build -o main .

# Etap uruchamiania Nginx z HTTP/3
FROM nginx:latest

# Instalacja openssl
RUN apt-get update && apt-get install -y openssl

COPY ./nginx.conf /etc/nginx/nginx.conf
COPY ./generate-ssl.sh /docker-entrypoint.d/

RUN chmod +x /docker-entrypoint.d/generate-ssl.sh
RUN /docker-entrypoint.d/generate-ssl.sh

COPY --from=builder /app/main /usr/share/nginx/html/main

EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]