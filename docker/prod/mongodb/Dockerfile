FROM alpine:3.9

COPY users.js /etc/
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
RUN chmod +x /usr/local/bin/docker-entrypoint.sh
COPY mongod.conf /etc/

RUN apk update &&\
    apk add mongodb mongodb-tools bash nano &&\
    mkdir -p /data/db

EXPOSE 27017

ENTRYPOINT "/usr/local/bin/docker-entrypoint.sh"