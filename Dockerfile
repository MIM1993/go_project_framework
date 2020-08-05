FROM daocloud.io/ubuntu:bionic

ENV MODULE gateway

ENV TZ 'Asia/Shanghai'
RUN echo $TZ > /etc/timezone && \
    apt-get update && DEBIAN_FRONTEND="noninteractive" apt-get install -y tzdata && \
    rm /etc/localtime && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    dpkg-reconfigure -f noninteractive tzdata && \
    apt-get clean

RUN mkdir -p /$MODULE

COPY output /$MODULE

WORKDIR /$MODULE

# ENV PATH "$PATH:/${MODULE}/bin"

