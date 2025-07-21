FROM golang:1.24.4-alpine AS build
WORKDIR /go/src/order-service

# Install git
RUN apk add --no-cache git

# Setup git configuration and credentials
RUN --mount=type=secret,id=git_credentials \
    git config --global credential.helper 'store --file=/root/.git-credentials' && \
    cp /run/secrets/git_credentials /root/.git-credentials && \
    chmod 600 /root/.git-credentials

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/order-service cmd/server/main.go

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/order-service /
EXPOSE 50053
CMD [ "/order-service" ]