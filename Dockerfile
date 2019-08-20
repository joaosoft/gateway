############################
# STEP 1 build binary
############################
FROM golang:alpine AS builder

LABEL maintainer="João Ribeiro <joaosoft@gmail.com>"

RUN apk update && apk add --no-cache \
	curl \
	mercurial \
	bash \
	dep \
	git

WORKDIR /go/src/gateway
COPY . .

RUN dep ensure

# build for raspberry pi 3
RUN GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=0 go build -o gateway ./main

RUN chmod +x gateway


############################
# STEP 2 run binary
############################
#FROM scratch
FROM alpine:latest

RUN apk update && apk --no-cache add \
	ca-certificates \
	curl \
	curl \
	bash

COPY --from=builder /go/src/gateway/gateway .
COPY ./config config

EXPOSE 8000
ENTRYPOINT ["./gateway"]