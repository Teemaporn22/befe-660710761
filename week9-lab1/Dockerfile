FROM golang:1.24.5 AS builder

WORKDIR /app

# เพื่อโหลด library ทั้งหมดที่ต้องใช้
COPY go.mod go.sum ./
RUN go mod download

# copy source code ทั้งหมด
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /root

COPY --from=builder /app/main .

ENTRYPOINT [ "./main" ]