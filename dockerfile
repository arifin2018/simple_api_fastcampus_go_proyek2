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

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.${OS}-${ARCH}.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

CMD ["sleep", "999999"]
