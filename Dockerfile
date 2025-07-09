FROM golang:1.24.4-alpine AS build

WORKDIR /go/src/order-service

COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/order-service cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/order-service /

EXPOSE 50053
CMD [ "/order-service" ]

