FROM golang:1.16

WORKDIR /go/src
ENV PATH=$PATH:/go/bin

RUN apt-get update && apt-get install -y \
  build-essential librdkafka-dev 

CMD ["tail", "-f", "/dev/null"]