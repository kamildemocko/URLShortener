FROM golang:latest AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o app ./cmd/app

RUN chmod +x app

FROM alpine:latest

WORKDIR /app

COPY .env .

COPY ./templates /app/templates

COPY --from=builder /app/app .

EXPOSE 80

CMD [ "/app/app" ]
