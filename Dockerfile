FROM golang

ADD . /go/src/github.com/PatipolThamdee/myproject
WORKDIR /go/src/github.com/PatipolThamdee/myproject
RUN go get ./...
RUN go build 
# RUN go install
# ENTRYPOINT ["/go/bin/myapp"]
#test comment eiei
EXPOSE 9090
        