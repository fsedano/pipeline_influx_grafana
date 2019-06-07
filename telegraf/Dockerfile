FROM ubuntu:18.04

RUN apt-get update && apt-get install -y vim wget gnupg iputils-ping influxdb-client
WORKDIR /app
RUN wget https://dl.influxdata.com/telegraf/releases/telegraf_1.11.0~rc1-1_amd64.deb
RUN apt install ./telegraf_1.11.0~rc1-1_amd64.deb
COPY data/config/telegraf.conf /etc/telegraf/

ENTRYPOINT [ "/bin/bash", "-c", "sleep 1000000" ]