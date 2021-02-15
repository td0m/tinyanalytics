FROM golang:1-alpine AS builder

WORKDIR /app

# download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy other files
COPY . .

# build
ENV CGO_ENABLED=0 GOOS=linux GO111MODULE=on
RUN go build -o main ./cmd/api/

# move binary to a minimal image
FROM scratch
COPY --from=builder /app/main /main

# run the binary
ENV PORT=80
EXPOSE 80
ENTRYPOINT ["/main"]