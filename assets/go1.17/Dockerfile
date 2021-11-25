FROM --platform=amd64 golang:1.17-alpine

ENV GOPATH /go
ENV GO111MODULE on
ENV CHOKIDAR_USEPOLLING true

# Ntp
RUN apk add --update --no-cache openntpd

# Bash Curl GIT
# - error: fork/exec
RUN apk add --update --no-cache bash curl git

# Make
RUN apk add --update --no-cache make gcc libc-dev

# Watcher
RUN go install github.com/cosmtrek/air@latest

# Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# GOsec
RUN go install github.com/securego/gosec/v2/cmd/gosec@latest

# Test
RUN go install github.com/vektra/mockery/v2@latest

# SSh agent
RUN \
    apk add --update --no-cache openssh && \
    mkdir -p /root/.ssh && chown -R root:root /root/.ssh && \
    echo "Host remotehost\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

CMD ["go", "--help"]
