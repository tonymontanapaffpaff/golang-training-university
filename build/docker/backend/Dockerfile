FROM golang:latest as build

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build ./cmd/main.go

FROM alpine:latest as production

COPY --from=build /backend/main .

EXPOSE 8080

CMD ["./main"]