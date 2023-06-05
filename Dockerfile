FROM golang:1.20.4

WORKDIR /workdir
COPY . . 

RUN go build -o prom-test main.go 

ENTRYPOINT [ "./prom-test" ]