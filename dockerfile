FROM golang:1.24

RUN apt-get update && apt-get install -y \
    default-mysql-client \
    iputils-ping \
    bash git openssh-client make curl tar

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV MIGRATE_VERSION=v4.15.2
ENV OS=linux
ENV ARCH=amd64

RUN go install golang.org/x/tools/cmd/goimports@latest && \
    go install golang.org/x/tools/gopls@latest
    
    # Install Go tools
RUN go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/...@v1.6.0 && \
    go install github.com/josharian/impl@v1.4.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest

ENV PATH="${PATH}:/root/go/bin"

RUN go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/cweill/gotests/...@v1.6.0 && \
    go install github.com/josharian/impl@v1.4.0 && \
    go install github.com/go-delve/delve/cmd/dlv@latest

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${OS}-${ARCH}.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

CMD ["sleep", "999999"]
