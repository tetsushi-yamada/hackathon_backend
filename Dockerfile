FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN go build -o ./cmd ./cmd/main.go

FROM gcr.io/distroless/base
WORKDIR /root
COPY --from=build /app/cmd/main .
CMD ["./main"]