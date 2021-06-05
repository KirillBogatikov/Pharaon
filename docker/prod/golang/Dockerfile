FROM golang:1.16-alpine

EXPOSE 80

RUN apk update && apk add git openssh &&\
    mkdir -p -m 0600 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts &&\
    git config --global url.git@github.com:.insteadOf https://github.com/ &&\
    go env -w GOPRIVATE=github.com/Projector-Solutions/* &&\
    go env -w GO111MODULE=on