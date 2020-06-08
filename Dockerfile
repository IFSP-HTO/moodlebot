FROM golang:1.14

WORKDIR /go/src/moodlebot

COPY . . 

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 25/tcp

CMD ["moodlebot"]