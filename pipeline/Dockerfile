FROM ubuntu:16.04
#RUN apt-get update -q --fix-missing
RUN mkdir /pipeline
COPY pipeline_data /pipeline
COPY json /data
COPY target/wait.sh /tmp
COPY target/id_rsa* /tmp/
#CMD ["/bin/bash", "-e", "/tmp/wait.sh"]
WORKDIR /running_data
CMD ["/bin/bash", "-c", "/pipeline/bin/pipeline -pem /tmp/id_rsa -config /pipeline/pipeline.conf_REWRITTEN > /tmp/salida.txt"]
