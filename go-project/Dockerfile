##FROM golang:1.18
FROM scratch
##ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/go-project
COPY . $GOPATH/src/github.com/go-project
##RUN go build -o go-project main.go

EXPOSE 8000
CMD ["./go-project"]