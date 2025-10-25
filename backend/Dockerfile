FROM golang:1.25.3-alpine3.22 AS builder

WORKDIR /app

COPY ./backend/go.mod .
COPY ./backend/go.sum .

ENV GOPROXY=https://goproxy.cn,direct
RUN go mod download

COPY ./backend .

RUN ls -l /app
RUN ls -l /app/main.go
RUN CGO_ENABLED=0 go build -ldflags='-s -w' -o /usr/local/bin/server ./main.go 

FROM alpine:3.22.2

EXPOSE 8080

COPY --from=builder /usr/local/bin/server /usr/local/bin/server
CMD ["/usr/local/bin/server"]