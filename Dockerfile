# get go
FROM golang:latest

# create the directories
RUN mkdir -p /go/src/github.com/hugocortes/paprika-api
RUN mkdir /go/src/github.com/hugocortes/paprika-api/modules
RUN mkdir /go/src/github.com/hugocortes/paprika-api/plugins

# set the work dir
WORKDIR /go/src/github.com/hugocortes/paprika-api

# copy necessary folders
COPY main.go /go/src/github.com/hugocortes/paprika-api
COPY modules/ /go/src/github.com/hugocortes/paprika-api/modules
COPY plugins /go/src/github.com/hugocortes/paprika-api/plugins

# configure
RUN go get ./
RUN go build -o paprika-api .
RUN touch .env

# run
CMD ["paprika-api"]
