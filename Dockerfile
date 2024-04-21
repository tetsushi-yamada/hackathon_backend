FROM --platform=linux/amd64 golang:1.22 as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd ./cmd/main.go

FROM gcr.io/distroless/base
WORKDIR /root
COPY --from=build /app/cmd/main .
EXPOSE 8080
CMD ["./main"]