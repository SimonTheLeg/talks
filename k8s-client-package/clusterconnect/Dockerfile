# Note: Build this image from within the 'k8s-client-package' folder and supply the Dockerfile via -f so
# that the go.mod and go.sum files are in the build context and can be copied into the image
# docker build --rm . -t clusterconnect -f clusterconnect/Dockerfile

FROM golang:1.11-alpine as builder

COPY go.mod go.sum /github.com/simontheleg/talks/k8s-client-package/
COPY clusterconnect/main.go /github.com/simontheleg/talks/k8s-client-package/clusterconnect/

WORKDIR /github.com/simontheleg/talks/k8s-client-package

# Install git as it is required by go modules
# And yep at the moment you also need gcc https://github.com/golang/go/issues/26988
RUN apk update && \
    apk add git gcc musl-dev && \
    cd clusterconnect && \
    go build


FROM alpine:3.9

COPY --from=builder /github.com/simontheleg/talks/k8s-client-package/clusterconnect/clusterconnect /

ENTRYPOINT [ "/clusterconnect" ]