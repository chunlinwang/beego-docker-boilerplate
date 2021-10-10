### PHP ###
ARG GOLANG_VERSION=1.17.2
ARG ALPINE_VERSION=3.14

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS base

RUN set -eux \
  && apk update && apk --no-cache add \
    postgresql-dev \
    libzip-dev \
    git \
	icu-libs \
	bash \
	autoconf \
	g++ \
	make \
	icu-dev \
	gcc \
	libtool \
	rabbitmq-c \
	rabbitmq-c-dev

# Init user
RUN addgroup -g 2000 -S app && adduser -G app -u 2000 -s /bin/bash -S app && sed -i "s/2000/2000/" /etc/group

# Configure Timezone
RUN ln -fs /usr/share/zoneinfo/Europe/Paris /etc/localtime && \
    echo "Europe/Paris" > /etc/timezone

RUN mkdir -p /var/app
WORKDIR /var/app

RUN chown app: -R /var/app

USER app

# Install beego & bee
RUN go get github.com/astaxie/beego
RUN go get github.com/beego/bee

#CMD ["bee", "run"]

EXPOSE 8081
