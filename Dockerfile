FROM golang:1.24.1-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR /go/cryptofox
COPY . .

RUN go install
RUN go build -o /go/cryptofox/bin/cryptofox

FROM scratch
COPY --from=builder /go/cryptofox/bin/cryptofox /go/bin/cryptofox
ENTRYPOINT ["/go/bin/cryptofox"]
