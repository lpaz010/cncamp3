FROM golang:1.18 AS build
WORKDIR /httpserver/
COPY --from=build /httpserver/httpserver /httpserver/httpserver 
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpserver main.go

FROM ubuntu
ENV MY_SERVICE_PORT=80
ENV MY_SERVICE_PORT1=80
LABEL multi.label1="value1" multi.label2="value2"
ADD bin/amd64/httpserver /httpserver
EXPOSE 80
WORKDIR /httpserver/
ENTRYPOINT ["./httpserver"]