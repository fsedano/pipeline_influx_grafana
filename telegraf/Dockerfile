FROM ubuntu:18.04
RUN apt-get update && apt-get install -y vim
RUN apt-get install -y wget
RUN apt-get install -y gnupg

RUN wget -qO- https://repos.influxdata.com/influxdb.key | apt-key add -
RUN echo "deb https://repos.influxdata.com/ubuntu bionic stable" | tee /etc/apt/sources.list.d/influxdb.list
RUN apt-get update && apt-get install -y telegraf
ENTRYPOINT [ "/bin/bash", "-c", "sleep 1000000" ]