FROM ubuntu:16.04
RUN apt-get update && apt-get install -y influxdb-client vim less
RUN mkdir /pipeline
COPY pipeline_data /pipeline
COPY json /data
WORKDIR /app
COPY target/bootstrap.sh /app
RUN chmod +x /app/bootstrap.sh
COPY target/id_rsa* /app/
WORKDIR /running_data
CMD ["/bin/bash", "-c", "/app/bootstrap.sh"]
