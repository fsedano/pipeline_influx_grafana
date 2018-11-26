This is a setup to perform real time visualisation and analysis of C9800 wireless controller telemetry data streaming.


![General diagram](diagrams/general.png?raw=true "General network diagram")


The setup is packaged in 3 different Docker containers, which are brought together using docker-compose.
Telemetry data from Cisco C9800 wireless controller is sent to [pipeline](https://developer.cisco.com/codeexchange/github/repo/cisco/bigmuddy-network-telemetry-pipeline/)

Pipeline is configured to accept telemetry data using grpc and output two streams:

Raw 'tap' data, for debugging purposes
[Influxdb](https://www.influxdata.com/time-series-platform/influxdb/) data output, filtered using the pipeline filter described on this guide.

Another container runs [Grafana](https://grafana.com/) which consumes Influxdb data to produce the real time visualization.

## Installation requirements
The only requirement to run the setup is to have docker and docker-compose installed

## Setup

On the Cisco C9800 wireless controller, configure telemetry data:

```
# TODO add telemetry config
```

Once Docker and Docker-compose are installed, just bring up docker-compose:

```
docker-compose up -d
```

The first the command is executed, the container for pipeline will be built, which will take some minutes. Eventually all 3 containers should come up:

```
Successfully tagged grafana_pipeline:latest
WARNING: Image for service pipeline was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Starting grafana_influx_1   ... done
Starting grafana_grafana_1  ... done
Creating grafana_pipeline_1 ... done
```

### Verification
To verify all 3 containers are up, you can use `docker-compose ps` to show the current status:

```
FSEDANOC-M-C1Z8:grafana fsedanoc$ docker-compose ps
       Name                     Command               State                       Ports
------------------------------------------------------------------------------------------------------------
grafana_grafana_1    /run.sh                          Up      0.0.0.0:3000->3000/tcp
grafana_influx_1     /entrypoint.sh influxd           Up      0.0.0.0:2003->2003/tcp, 0.0.0.0:8086->8086/tcp
grafana_pipeline_1   /bin/bash -c /pipeline/bin ...   Up      0.0.0.0:57500->57500/tcp
```

## Accessing each container

To access each of the container, from the main repository path, use `docker-compose exec <container_name> bash`, where `<container-name>` is either `grafana`, `influx` or `pipeline`

## Real time logging of raw data

Pipeline is configured to dump the raw data from the controller to the file `/data/dump.txt`, and the data after filtering to `/running_data/metricsdump.txt*`.
Please note this configuration is not appropiate for a production environment, since the raw data files will grow quickly. Data dump can be disabled by changing pipeline configuration, as described on the following section.

## Updating pipeline configuration

The entire pipeline, including it's config file is in the `pipeline` container, which is built first time docker-compose is brought up, or when `docker-compose build` is executed.

To change pipeline configuration, from your repository, edit the file `pipeline/pipeline_data/pipeline.conf_REWRITTEN`

Once that file is update, stop and rebuild the pipeline container with

```
docker-compose down pipeline
docker-compose up -d --build pipeline
```

## Pipeline filters
When pipeline receives gRPC data it runs it thru a filter, which is described via the JSON file `/data/metrics.json`. This file needs to match the data being received from the wireless controller.

