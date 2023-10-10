FROM ubuntu:latest

RUN  apt-get update
RUN apt-get upgrade -y
RUN apt-get install -y sudo 

RUN apt-get install -y automake 
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y gtk-doc-tools 
RUN apt-get install autoconf 
RUN apt-get install -y gobject-introspectio
RUN apt-get install -y libglib2.0-dev

RUN apt-get install -y vim 
RUN  apt-get install -y tree 
RUN apt-get install -y git

RUN apt-get install -y libvips-dev

RUN  apt-get install -y software-properties-common 
RUN add-apt-repository -y ppa:longsleep/golang-backports 
RUN apt-get install -y golang

WORKDIR /go/work/storagesvc
COPY go.mod go.sum  /go/work/storagesvc/
RUN go mod download




CMD ["/bin/bash"]