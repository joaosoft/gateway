############################
# STEP 1 build binary
############################
FROM golang:alpine AS builder

LABEL maintainer="João Ribeiro <joaosoft@gmail.com>"

RUN apk update && apk add --no-cache \
	curl \
	bash \
	dep \
	git

WORKDIR /go/src/gateway
COPY . .

RUN dep ensure

RUN GOOS=linux GOARCH=arm GOARM=5 CGO_ENABLED=0 go build -a -installsuffix cgo -o gateway .

RUN chmod +x gateway


############################
# STEP 2 run binary
############################
FROM scratch
COPY --from=builder /go/src/gateway/gateway .

EXPOSE 8000
ENTRYPOINT ["./gateway"]