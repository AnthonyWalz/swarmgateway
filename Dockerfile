FROM golang:1.8-alpine3.6

# install git
RUN apk --update add \
	bind-tools \
	&& rm /var/cache/apk/*

WORKDIR /go/src

ADD . /go/src

CMD ["go", "run", "main.go"]

EXPOSE 80