FROM ubuntu:plucky

RUN echo "deb https://ppa.launchpadcontent.net/longsleep/golang-backports/ubuntu ubuntu-plucky main" | tee /etc/apt/sources.list.d/docker.list
RUN apt update

ENV GO_VERSION=1.24

RUN apt install -y golang-${GO_VERSION}
RUN ln -s "/usr/lib/go-${GO_VERSION}/bin/go" /usr/local/bin/

RUN apt-get clean && apt-get update && apt-get install -y locales wget
RUN locale-gen en_US.UTF-8
RUN update-locale LC_ALL=en_US.UTF-8 LANG=en_US.UTF-8
ENV LANG=en_US.UTF-8
ENV LC_CTYPE="en_US.UTF-8"
ENV LC_NUMERIC="en_US.UTF-8"
ENV LC_TIME="en_US.UTF-8"
ENV LC_COLLATE="en_US.UTF-8"
ENV LC_MONETARY="en_US.UTF-8"
ENV LC_MESSAGES="en_US.UTF-8"
ENV LC_PAPER="en_US.UTF-8"
ENV LC_NAME="en_US.UTF-8"
ENV LC_ADDRESS="en_US.UTF-8"
ENV LC_TELEPHONE="en_US.UTF-8"
ENV LC_MEASUREMENT="en_US.UTF-8"
ENV LC_IDENTIFICATION="en_US.UTF-8"
ENV LC_ALL=
   
RUN apt install -y prusa-slicer

WORKDIR /build

COPY ../ .

RUN go build -o /bin/slic3r-api ./cmd

RUN rm -rf /build/*
 
CMD [ "/bin/slic3r-api" ]