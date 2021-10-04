FROM ubuntu:20.04

RUN apt-get update && \
apt-get install curl ffmpeg jq language-pack-ja -y && \
apt-get clean

ENV LC_ALL=ja_JP.UTF-8

EXPOSE 80

COPY src/ /usr/local/bin/
RUN chmod +x /usr/local/bin/* && \
mkdir /convert

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]