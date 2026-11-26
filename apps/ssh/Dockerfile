# Build a static Wish SSH server. This is a TCP service, not an HTTP Worker.
FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /ssh-portfolio ./cmd/ssh-portfolio

FROM alpine:3.20
RUN adduser -D -H -u 65532 portfolio \
	&& mkdir -p /data \
	&& chown portfolio:portfolio /data
COPY --from=build /ssh-portfolio /usr/local/bin/ssh-portfolio
ENV SSH_ADDR=0.0.0.0:2222 \
	HOST_KEY_PATH=/data/host_ed25519
USER portfolio
VOLUME ["/data"]
EXPOSE 2222
ENTRYPOINT ["/usr/local/bin/ssh-portfolio"]
