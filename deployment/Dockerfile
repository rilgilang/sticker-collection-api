# STEP 1: Build the application
FROM --platform=linux/amd64 golang:1.20.6-alpine3.17 as builder

ARG GO_BUILD_COMMAND="go build -tags static_all"

RUN apk update && apk --no-cache add build-base git bash coreutils openssh openssl

WORKDIR /go/src/github.com/rilgilang/sticker-collection-api

COPY . .

RUN go mod tidy && go mod download && go mod vendor
RUN eval $GO_BUILD_COMMAND

# STEP 2: Build the runtime image
FROM --platform=linux/amd64 alpine:3.17.0

RUN apk --no-cache add tzdata && \
    ln -snf /usr/share/zoneinfo/Asia/Jakarta /etc/localtime && echo "Asia/Jakarta" > /etc/timezone

ENV PROJECT_DIR=/opt/github.com/rilgilang/sticker-collection-api
ARG BUILD_VERSION
ENV DD_VERSION=$BUILD_VERSION

RUN mkdir -p $PROJECT_DIR

WORKDIR $PROJECT_DIR

COPY --from=builder /go/src/github.com/rilgilang/sticker-collection-api/sticker-collection-api .

RUN chmod +x sticker-collection-api

CMD ["/opt/github.com/rilgilang/sticker-collection-api/sticker-collection-api"]
