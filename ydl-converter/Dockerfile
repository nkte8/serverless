FROM ubuntu:20.04

ENV TZ=Asia/Tokyo
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get install curl ffmpeg jq language-pack-ja -y && \
    apt-get autoremove && \ 
    apt-get -y autoclean && \ 
    apt-get clean && \ 
    rm -rf /var/lib/apt/lists/*

ENV LC_ALL=ja_JP.UTF-8

EXPOSE 80

COPY src/ /usr/local/bin/
RUN chmod +x -R /usr/local/bin

WORKDIR /convert
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]