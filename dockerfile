# ---------- frontend build ----------
FROM node:24-alpine AS frontend-builder

WORKDIR /app/frontend

COPY frontend/package*.json ./
RUN npm ci

COPY frontend/ ./
RUN npm run build


# ---------- backend build ----------
FROM golang:1.26-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# переносим только результат сборки фронта в фиксированное место
COPY --from=frontend-builder /app/frontend/dist ./dist

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o client-service ./cmd


# ---------- runtime ----------
FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=backend-builder /app/client-service ./
COPY --from=backend-builder /app/dist ./dist

EXPOSE 8080

CMD ["./client-service"]