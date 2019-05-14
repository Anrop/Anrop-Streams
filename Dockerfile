FROM golang:1.12-alpine AS builder
RUN apk --no-cache add git
RUN go get github.com/constabulary/gb/...
WORKDIR /build/
COPY src/ /build/src/
COPY vendor/ /build/vendor/
RUN gb build all

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /build/bin/streams /app/streams
CMD /app/streams
