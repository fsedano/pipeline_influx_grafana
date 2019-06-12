# docker run -ti -v `pwd`:/runapp -p 4000:4000 fsedano/rest-client sh
FROM ubuntu:18.04
RUN apt-get update && apt-get install -y python3 python3-pip wget

RUN pip3 install flask flask_restful flask_cors requests netmiko
RUN apt-get install -y git
WORKDIR /app
RUN git clone https://github.com/networktocode/ntc-templates.git
ENV NET_TEXTFSM=/app/ntc-templates/templates/
COPY target/rest-server.py /app

ENTRYPOINT [ "python3", "rest-server.py" ]
#ENTRYPOINT [ "/bin/sh", "-c", "sleep 100000" ]

