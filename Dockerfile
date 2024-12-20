FROM golang:1.23.2-bullseye AS base

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid 65532 \
  zomboid-server-manager

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go mod verify

# Copy the source from the current directory to the Working Directory inside the container
COPY ./cmd /app/cmd
COPY ./internal /app/internal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /zomboid-server-manager ./cmd

FROM scratch

COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

COPY --from=base /zomboid-server-manager .

COPY ./test-data /test-data

USER zomboid-server-manager:zomboid-server-manager

CMD ["./zomboid-server-manager"]