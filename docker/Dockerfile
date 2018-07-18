FROM golang:alpine
MAINTAINER quqi99 <quqi99@163.com>
RUN apk update && \
    apk upgrade && \
    apk add git bash 
RUN go get -u github.com/zhhuabj/ss/...
COPY ss.json /root/.ss.json
RUN rm -rf /go/src/github.com
CMD ["/go/bin/ssserver"]
