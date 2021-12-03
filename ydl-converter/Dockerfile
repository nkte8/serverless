FROM alpine:3.14.3

ENV TZ=Asia/Tokyo
ENV DEBIAN_FRONTEND=noninteractive

RUN apk --no-cache --update add curl ffmpeg jq bash
ENV LC_ALL=ja_JP.UTF-8

EXPOSE 80

COPY src/ /usr/local/bin/
RUN chmod +x -R /usr/local/bin

WORKDIR /convert
ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]