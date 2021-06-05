FROM 192.168.5.1:49004/pharaon/golang:latest

EXPOSE 80

WORKDIR /go/src/pharaon-migration
COPY . .

RUN eval `ssh-agent` &&\
    chmod 0600 github &&\
    ssh-add github &&\
    go get && go build &&\
    rm -f github

ENTRYPOINT pharaon-migration