############################
# STEP 1 build binary
############################
FROM golang:alpine AS builder

LABEL maintainer="João Ribeiro <joaosoft@gmail.com>"

RUN apk update && apk add --no-cache \
	dep \
	git

WORKDIR /go/src/gateway
COPY . .

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway .



############################
# STEP 2 run binary
############################
FROM scratch
COPY --from=builder /go/src/gateway/gateway .

ENTRYPOINT ./gateway