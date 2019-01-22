This is a setup to perform real time visualisation and analysis of C9800 wireless controller telemetry data streaming. It uses [dial-out](https://xrdocs.io/telemetry/blogs/2017-01-20-model-driven-telemetry-dial-in-or-dial-out/) telemetry to push a periodic stream of wireless operational data to an open source collector: [pipeline](https://developer.cisco.com/codeexchange/github/repo/cisco/bigmuddy-network-telemetry-pipeline/)


![General diagram](diagrams/general.png?raw=true "General network diagram")

A number of different flows are possible, since pipeline is capable to export to a great variety of sources. On this particular demo, InfluxDB  + Grafana are used to display live data

![Pipeline diagram](diagrams/pipeline.png?raw=true "Pipeline")



### Structure of docker composition

The setup is packaged in 4 different Docker containers, which are brought together using docker-compose.

![Composite diagram](diagrams/compose-internal1.png?raw=true "Composite diagram")


Telemetry data from Cisco C9800 wireless controller is sent to [pipeline](https://developer.cisco.com/codeexchange/github/repo/cisco/bigmuddy-network-telemetry-pipeline/)

Pipeline is configured to accept telemetry data using grpc and output two streams:

Raw 'tap' data, for debugging purposes
[Influxdb](https://www.influxdata.com/time-series-platform/influxdb/) data output, filtered using the pipeline filter described on this guide.

Another container runs [Grafana](https://grafana.com/) which consumes Influxdb data to produce the real time visualization.

Also, a container for yang-explorer is present, so we can interactively browse the YANG model of the device. 
In order to use the yang-explorer UI, we need to point our browser to the IP address of the docker-compose VM, port 8088, with a flash-enabled browser.

## Installation requirements
The only requirement to run the setup is to have docker and docker-compose installed

## Setup

On the Cisco C9800 wireless controller, configure telemetry data:

```
telemetry ietf subscription 0
 encoding encode-kvgpb
 filter xpath /wireless-client-oper:client-oper-data/traffic-stats
 source-address 9.9.71.50
 stream yang-push
 update-policy periodic 1000
 receiver ip address 9.9.71.130 57500 protocol grpc-tcp
 
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
docker-compose stop pipeline
docker-compose up -d --build pipeline
```

## Pipeline filters
When pipeline receives gRPC data it runs it thru a filter, which is described via the JSON file `/data/metrics.json`. This file needs to match the data being received from the wireless controller.

 On this example, we'll be plotting client stats (traffic, RSSI). The JSON filter needs to match the YANG data model, and is as follows:

 ```
[
	{
		"basepath" : "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/radio-oper-data",
		"spec" : {
			"fields" : [
				{"name" : "radio-slot-id", "tag" : true},
				{"name" : "radio-type"},
				{"name" : "admin-state", "track" : true},
				{"name" : "wtp-mac", "tag" : true},
				{"name" : "phy-ht-cfg", 
					"fields" : [
						{"name" : "cfg-data", 
							"fields" : [
								{"name" : "chan-width"},
								{"name" : "curr-freq"}
							]
						}
					]
				}
			]
		}
	},
   {
           "basepath" : "Cisco-IOS-XE-wireless-access-point-oper:access-point-oper-data/ethernet-if-stats",
           "spec" : {
                   "fields" : [
                           {"name" : "admin-state"},
                           {"name" : "if-name", "tag" : true},
                           {"name" : "rx-total-bytes"}
                   ]
           }
   },
   {
               "basepath" : "Cisco-IOS-XE-wireless-client-oper:client-oper-data/traffic-stats",
               "spec" : {
                       "fields" : [
                               {"name" : "ms-mac-address", "tag": true},
                               {"name": "bytes-rx"},
                				{"name": "bytes-tx"},
                				{"name": "data-retries"},
                				{"name": "mic-mismatch"},
                				{"name": "mic-missing"},
                				{"name": "most-recent-rssi"},
                				{"name": "most-recent-snr"},
                				{"name": "pkts-rx"},
                				{"name": "pkts-tx"}
                       ]
               }
       }
]

 ```

## Creating graph
Once the containers are brought up, Grafana web interface can be accessed via port 3000 on the instance public IP.

Use `admin / admin` as the initial user and password.


### Datasource definition

Setup a data source as shown on the following capture, the important parameters are:
- Type: influxdb
- URL: http://influx:8086
- Access: Server
- Database: vwlcdata

![Data source configuration](diagrams/datasource.png?raw=true "Data source")

### Graph definition

Once datasource is setup, you can define graphs using the interactive builder. An example can be found below:

![Graph definition](diagrams/graph-definition-1.png?raw=true "Graph definition")

This is the resulting graph output:

![Client RSSI Graph](diagrams/graph-example-rssi.png?raw=true "Client RSSI graph")

## Changing controller config data from Dashboard

It's also possible to change the controller configuration programatically from the dashboard. The buttons on this example dashboard will trigger a config change to the controller via RESTconf

![Buttons](diagrams/dashboard-buttons.png?raw=true "Buttons")