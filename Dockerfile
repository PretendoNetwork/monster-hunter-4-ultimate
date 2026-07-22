# --- builder ---
FROM golang:1.25-alpine3.22 as builder
LABEL stage=builder
RUN apk add git
WORKDIR /build

COPY go.* ./
RUN go mod download

COPY . ./
ARG BUILD_STRING=pretendo.mh4u.docker
RUN go build -ldflags "-X 'main.serverBuildString=${BUILD_STRING}'" -v -o server

# --- runner ---
FROM alpine:3.22 as runner
WORKDIR /build

COPY --from=builder /build/server /build/
CMD ["/build/server"]
