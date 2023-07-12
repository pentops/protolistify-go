FROM golang:1.20 AS builder

WORKDIR /src

ADD . .
ARG VERSION
RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -ldflags="-X main.Version=$VERSION" \
    -o /protoc-gen-listify .

FROM scratch
COPY --from=builder /protoc-gen-listify /
ENTRYPOINT ["/protoc-gen-listify"]
