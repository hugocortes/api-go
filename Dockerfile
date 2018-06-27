# get go
FROM golang:latest

# create the directories
RUN mkdir -p /go/src/github.com/hugocortes/api-go
RUN mkdir /go/src/github.com/hugocortes/api-go/modules
RUN mkdir /go/src/github.com/hugocortes/api-go/plugins

# set the work dir
WORKDIR /go/src/github.com/hugocortes/api-go

# copy necessary folders
COPY main.go /go/src/github.com/hugocortes/api-go
COPY modules/ /go/src/github.com/hugocortes/api-go/modules
COPY plugins /go/src/github.com/hugocortes/api-go/plugins

# configure
RUN go get ./
RUN go build -o api-go .
RUN touch .env

# run
CMD ["api-go"]
