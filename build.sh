#!/bin/bash

# pre-flight syntax validation
go fmt ./...

# disable cgo, target linux
CGO_ENABLE=0 
GOOS=linux 
go build \
  # specify netgo instead of system
  -tags netgo \ 
  # omit debug to quash size
  -ldflags '-w -s' \
  # verbose, rebuild
  -v -a \
  # use the generic binary name 'app'
  -o app

# compress! 
upx -7 -qq app && \
  upx -t app && \
  mv ./app /go/bin/app