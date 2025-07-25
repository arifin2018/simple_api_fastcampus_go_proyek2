FROM golang:1.24-alpine

# Install development tools and dependencies
RUN apk update && apk add --no-cache \
    git \
    curl \
    bash

WORKDIR /app

# Copy go.mod first
COPY go.mod ./

# Initialize go.mod if main.go is empty or doesn't exist
RUN go mod tidy

# Copy all project files
COPY . .

ENV MIGRATE_VERSION=v4.15.2
ENV OS=linux
ENV ARCH=amd64
ENV PATH="${PATH}:/root/go/bin"

RUN go install golang.org/x/tools/cmd/goimports@latest && \
    go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/...@v1.6.0 && \
    go install github.com/josharian/impl@v1.4.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest


RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${OS}-${ARCH}.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

# For dev container, keep running
CMD ["sleep", "infinity"]
