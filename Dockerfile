FROM alpine:latest

EXPOSE 8080

# The command to run
CMD ["/go-test-kit"]

ARG BUILD_TAG=unknown
LABEL BUILD_TAG=$BUILD_TAG

COPY bin/go-test-kit /go-test-kit

