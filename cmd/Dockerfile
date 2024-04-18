FROM golang:1.22 as build

WORKDIR /app
COPY . .
RUN go build -tags netgo -installsuffix netgo -o main .

FROM gcr.io/distroless/base
WORKDIR /root
COPY --from=build /app/main .
CMD ["./main"]