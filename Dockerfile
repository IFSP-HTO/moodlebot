FROM golang:1.14

WORKDIR /go/src/moodlebot

COPY . . 

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 587/tcp

CMD ["moodlebot"]