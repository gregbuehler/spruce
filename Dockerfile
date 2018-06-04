# BUILD ----------------------------------------------------------------------
# build in language container
FROM golang:alpine as build

# install build dependencies (e.g. upx)
RUN apk add --no-cache git build-base && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk add --no-cache upx

# configure workspace
WORKDIR $GOPATH/src/github.com/gregbuehler/spruce
COPY . .

# resolve deps
RUN go get -u github.com/golang/dep/...
RUN dep ensure -v -vendor-only

# syntax check, build & compress static bin
RUN build.sh


# BUNDLE ---------------------------------------------------------------------
# bundle and execute in distroless
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]