FROM golang:1.24.4-alpine AS build

WORKDIR /go/src/order-service

COPY go.mod go.sum ./
RUN apk add --no-cache git

ARG GITHUB_TOKEN
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go env -w GOPRIVATE=github.com/commerce-app-demo/* 

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/order-service cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/order-service /

EXPOSE 50053
CMD [ "/order-service" ]

