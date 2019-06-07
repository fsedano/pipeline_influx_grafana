FROM ubuntu:18.04
ENV HTTP_PROXY http://proxy.esl.cisco.com:80
ENV HTTPS_PROXY http://proxy.esl.cisco.com:80
ENV http_proxy http://proxy.esl.cisco.com:80
ENV https_proxy http://proxy.esl.cisco.com:80
ENV no_proxy ".cisco.com"

RUN apt-get update && apt-get install -y vim
RUN apt-get install -y wget
RUN apt-get install -y gnupg
RUN apt-get install -y iputils-ping influxdb-client
WORKDIR /app
RUN wget https://dl.influxdata.com/telegraf/releases/telegraf_1.11.0~rc1-1_amd64.deb
RUN apt install ./telegraf_1.11.0~rc1-1_amd64.deb
#RUN wget -qO- https://repos.influxdata.com/influxdb.key | apt-key add -
#RUN echo "deb https://repos.influxdata.com/ubuntu bionic stable" | tee /etc/apt/sources.list.d/influxdb.list
#RUN apt-get update && apt-get install -y telegraf
COPY data/config/telegraf.conf /etc/telegraf/

ENTRYPOINT [ "/bin/bash", "-c", "sleep 1000000" ]