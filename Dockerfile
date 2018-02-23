FROM debian

RUN apt-get update -y && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates

ADD main /
CMD ["/main"]
