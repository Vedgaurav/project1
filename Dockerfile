FROM golang:alpine

EXPOSE 8080

WORKDIR /go/src/project1

COPY . .

CMD go run travel-main.go

RUN go install  project1