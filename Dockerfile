FROM golang:1.23.0-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext

COPY ["go.mod", "./"]
RUN go mod download

COPY . ./
RUN go build -o app test/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src /

CMD ["/app"]
