FROM golang:alpine AS builder

COPY . /app

WORKDIR /app

RUN go build -ldflags "-s -w" -o /app/dolistasafado ./cmd/dolistasafado

FROM alpine:latest

RUN apk --no-cache add bash ca-certificates

COPY --from=builder /app/dolistasafado /dolistasafado

EXPOSE 80

ENTRYPOINT [ "/dolistasafado" ]