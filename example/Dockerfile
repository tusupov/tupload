FROM centos:7

WORKDIR /root

RUN yum install -y wget && \
    wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz

RUN tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz && \
    mkdir go && mkdir go/src && mkdir go/bin && mkdir go/pkg && \
    mkdir go/src/upload

ENV PATH=${PATH}:/usr/local/go/bin GOROOT=/usr/local/go GOPATH=/root/go

RUN go get github.com/tusupov/tupload

ADD src go/src
ADD main.go go/src/upload

RUN go build upload && go install upload

EXPOSE 80

CMD ./go/bin/upload
