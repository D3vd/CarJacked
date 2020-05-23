FROM golang:1.13.5

WORKDIR /go/src/github.com/R3l3ntl3ss/CarJacked

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080

CMD [ "go", "run", "." ]