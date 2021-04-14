FROM golang:alpine AS builder


ENV GO111MODULE=ON \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o bubble .

FROM scrath

COPY --form=builder /build/bubble /

EXPOSE 8888

ENTRYPOINT ["/bubble"]
