FROM golang:1.9.4-alpine3.7
MAINTAINER Ajay<ajay@plivo.com>

RUN apk --no-cache add --virtual .build-deps bash postgresql-dev gcc musl-dev git glide openssh

# to pip install from private repo
RUN mkdir -p /root/.ssh
# RUN ls -al /root
COPY .ssh /root/.ssh
RUN chmod 400 /root/.ssh/id_rsa && \
	chown -R root:root /root/.ssh/

WORKDIR /opt/app
COPY . /opt/app
ENV GOPATH /opt/app
RUN echo "$PWD"
RUN echo "$GOPATH"

WORKDIR /opt/app/src
RUN glide install
# RUN echo "$PWD"
# RUN ls -l .
# RUN echo "$PWD"


RUN	apk del .build-deps && \
	rm -rf /root/.ssh/

WORKDIR /opt/app
RUN source ./src/config/env.sh

EXPOSE 8090

CMD ["go", "run", "src/main.go", "-e", "PROD"]
