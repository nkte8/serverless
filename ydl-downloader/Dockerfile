FROM python:3.8.12-alpine3.14 as builder

WORKDIR /workspace
ENV DEBIAN_FRONTEND=noninteractive
RUN apk --no-cache --update add git gcc musl-dev

RUN pip3 install --no-cache --upgrade git+https://github.com/yt-dlp/yt-dlp.git@release && \
    pip3 install --no-cache --upgrade requests pyinstaller

COPY src/ .
RUN pyinstaller ./ydl-downloader.py --onefile

FROM alpine:3.14.3
WORKDIR /download
ENV TZ=Asia/Tokyo
COPY --from=builder /workspace/dist/ydl-downloader /usr/local/bin/
ENTRYPOINT ["ydl-downloader"]