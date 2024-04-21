ARG GO_VERSION

FROM golang:${GO_VERSION}

WORKDIR /usr/src/appstore-notify-sample

COPY go.mod .
COPY go.sum .
RUN go mod download

RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go

WORKDIR ${GOPATH}
