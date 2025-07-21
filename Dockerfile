FROM golang:1.24.4-alpine AS build

WORKDIR /go/src/order-service

# Add ARG for GITHUB_TOKEN
ARG GITHUB_TOKEN

RUN apk add --no-cache git

# Configure Git for private module access
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/order-service cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/order-service /

EXPOSE 50053
CMD [ "/order-service" ]

