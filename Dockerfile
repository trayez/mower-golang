### BUILDER
FROM golang:1.12.7-alpine as builder
ENV GOPATH="/mower"


WORKDIR /mower

ADD . /mower/

RUN cd /mower
RUN go build -o /mower/bin/mower src/exo/main.go
#RUN cd /mower && go test exo/mower exo/orientation exo/parser

### RUNTIME
FROM golang:1.12.7-alpine 

WORKDIR /usr/bin

COPY --from=builder /mower/bin/mower /usr/bin/
