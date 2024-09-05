# -- multistage docker build: stage #1: build stage
FROM golang:1.22.0-alpine AS build

RUN mkdir -p /go/src/github.com/AnumaNetwork/anumad

WORKDIR /go/src/github.com/AnumaNetwork/anumad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache anumad dependencies
RUN go mod download

COPY . .
RUN mkdir -p /anuma/bin/
RUN go build $FLAGS -o /anuma/bin/ ./cmd/...

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /root/

RUN apk add --no-cache ca-certificates tini

COPY --from=build /anuma/bin/* /usr/bin/

ENTRYPOINT [ "/usr/bin/anumad" ]
