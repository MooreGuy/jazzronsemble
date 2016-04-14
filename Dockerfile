FROM golang

ADD . /go/src/github.com/jazzronsemble
ADD app /srv/jazzronsemble
RUN go get github.com/jazzronsemble && go install github.com/jazzronsemble
ENTRYPOINT /go/bin/jazzronsemble

EXPOSE 8080
