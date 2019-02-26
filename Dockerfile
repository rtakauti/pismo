FROM golang

RUN apt-get update
RUN apt-get install git

WORKDIR /go/src/app
COPY ./go/*.go ./

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80
CMD ["app"]