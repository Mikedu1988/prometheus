FROM golang:1.13 as builder
LABEL maintainer="Mike Du <duliyang@juzhou.io>"
WORKDIR /go/src/app
COPY . .

#use go center as go proxy cloud result in package difference
RUN GOPROXY=https://gocenter.io go build -o server cmd/server/main.go
#CMD ["./server"]


FROM centos:7

WORKDIR /root/

COPY --from=builder /go/src/app/server .
RUN chmod u+x server
EXPOSE 9909
CMD ["./server"]