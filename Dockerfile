FROM golang:1.22.2 AS base
WORKDIR /usr/src/tech-interview-scrappers
COPY . .

FROM base AS development
WORKDIR /usr/src/tech-interview-scrappers
COPY . .
RUN go install github.com/air-verse/air@latest
EXPOSE 3001
ENTRYPOINT ["air"]

FROM base AS builder
WORKDIR /usr/src/tech-interview-scrappers
COPY . .
RUN go build -a --installsuffix cgo --ldflags="-s" -o main

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=builder /usr/src/tech-interview-scrappers .
EXPOSE 3001
ENTRYPOINT ["./main"]