FROM golang:latest as build

WORKDIR /k3s-task2

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/main.go

FROM alpine:latest as production

COPY --from=build /k3s-task2/main .

EXPOSE 8080

CMD ["./main"]