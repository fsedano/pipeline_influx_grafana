FROM ubuntu:16.04
RUN mkdir /pipeline
COPY pipeline_data /pipeline
COPY json /data
COPY target/wait.sh /tmp
COPY target/id_rsa* /tmp/
WORKDIR /running_data
CMD ["/bin/bash", "-c", "/pipeline/bin/pipeline -pem /tmp/id_rsa -config /pipeline/pipeline.conf_REWRITTEN > /tmp/salida.txt"]
