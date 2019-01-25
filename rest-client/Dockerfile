# docker run -ti -v `pwd`:/runapp -p 4000:4000 fsedano/rest-client sh
FROM alpine:3.8

RUN apk add --no-cache python3 && \
    python3 -m ensurepip && \
    rm -r /usr/lib/python*/ensurepip && \
    pip3 install --upgrade pip setuptools && \
    if [ ! -e /usr/bin/pip ]; then ln -s pip3 /usr/bin/pip ; fi && \
    if [[ ! -e /usr/bin/python ]]; then ln -sf /usr/bin/python3 /usr/bin/python; fi && \
    rm -r /root/.cache

RUN pip install flask flask_restful flask_cors requests
WORKDIR /app
COPY target/rest-server.py /app
ENTRYPOINT [ "python", "rest-server.py" ]

