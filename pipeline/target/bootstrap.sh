#!/bin/bash
false
while [ $? -ne 0 ]; do
    echo "Creating vwlc data DB"
    sleep 1
    influx -host influx -execute 'create database vwlcdata'
done

/pipeline/bin/pipeline -pem /app/id_rsa -config /pipeline/pipeline.conf_REWRITTEN