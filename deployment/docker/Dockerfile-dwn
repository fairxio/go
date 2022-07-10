FROM golang:latest
MAINTAINER Michael Dowling

RUN mkdir -p /opt/fairx && \
    mkdir -p /etc/fairx/src

COPY . /opt/fairx/src/

RUN cd /opt/fairx/src && go mod download
RUN cd /opt/fairx/src/applications/fairx && go build && go install
RUN rm -Rf /opt/fairx/src

VOLUME /etc/fairx
VOLUME /opt/fairx

CMD ["fairx", "dwn"]

