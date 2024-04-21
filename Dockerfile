FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN go build --platform linux/amd64 -o ./cmd ./cmd/main.go

FROM gcr.io/distroless/base
WORKDIR /root
COPY --from=build /app/cmd/main .
EXPOSE 8080
CMD ["./main"]