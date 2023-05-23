#
# STAGE 1: build
#
FROM golang:1.15-alpine3.12 as build
 
WORKDIR /source
COPY . .
 
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/main *.go
 
#
# STAGE 2: run
#
FROM alpine:3.12 as run
 
COPY --from=build /source/bin/main /main
 
ENTRYPOINT ["/main"]