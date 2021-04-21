FROM golang:1.16

ARG PACKAGE

ENV LISTEN=8000
ENV package_name=$PACKAGE

WORKDIR /healthy

COPY . /healthy
RUN go build ./cmd/${package_name}

CMD ./$package_name --http-port=$LISTEN
