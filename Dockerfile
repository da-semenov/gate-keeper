FROM golang:1.22.2

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o gate-keeper ./main.go

CMD ["./gate-keeper"]