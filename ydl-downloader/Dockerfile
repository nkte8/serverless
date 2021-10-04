FROM ubuntu:20.04

RUN apt-get update && \
apt-get install python3 python3-pip git language-pack-ja -y && \
apt-get clean

ENV LC_ALL=ja_JP.UTF-8

RUN pip3 install --upgrade git+https://github.com/yt-dlp/yt-dlp.git@release && \
pip3 install --upgrade requests 

EXPOSE 80

COPY src/ /usr/local/bin/
RUN chmod +x /usr/local/bin/* && \
mkdir /download

WORKDIR /download

ENTRYPOINT ["python3","/usr/local/bin/main.py"]