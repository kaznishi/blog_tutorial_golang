FROM golang:1-stretch

RUN go get -v github.com/tockins/realize \
    && go get -v github.com/golang/dep/cmd/dep

COPY . /go/src/github.com/kaznishi/blog_tutorial_golang
WORKDIR /go/src/github.com/kaznishi/blog_tutorial_golang

CMD [ "realize", "start"]