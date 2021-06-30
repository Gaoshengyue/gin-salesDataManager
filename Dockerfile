FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/dolphin/salesManager
COPY . $GOPATH/src/dolphin/salesManager
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./salesManager"]
