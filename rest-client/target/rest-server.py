from flask import Flask
from flask_restful import Resource, Api, reqparse
from flask_cors import CORS
from netmiko import ConnectHandler
import argparse
import requests, json
import sys
import logging

app = Flask(__name__)
CORS(app)
api = Api(app)



def sanitize(sanitized_output, output):
    #print(f"Sanintize in: {output}")
    if len(output) > 0:
        try:
            info = output[0].items()
        except:
            print(f"Can't parse: {output}")
            return sanitized_output

    for k,v in info:
        try:
            sanitized_output[k] = int(v)
        except:
            sanitized_output[k] = v
    #print(f"Sanintize out: {sanitized_output}")
    return sanitized_output

def count_radios_up(output, sanitized_output):
    #print(f"Processing {output}")
    radio_up_count = 0
    radio_down_count = 0
    for radio in output:
        try:
            if radio['slot'] == '1':
                if radio['oper_state'] == 'Up':
                    radio_up_count = radio_up_count + 1
                else:
                    radio_down_count = radio_down_count + 1
        except:
            pass
    sanitized_output['radios_up'] = radio_up_count
    sanitized_output['radios_down'] = radio_down_count
    return sanitized_output

def compute_fabric_state(output, sanitized_output):
    for node in output:
        if node['fabric_state'] == 'Up':
            fabric_value = 1
        else:
            fabric_value = 0
        sanitized_output[f"fabric_state_{node['fabric_node_ip']}"] = fabric_value

    return sanitized_output

class GetCLI(Resource):
    def get(self):
        net_connect = ConnectHandler(**cisco)
        
        print("Starting request")        
        sanitized_output = {}
        output = net_connect.send_command("show radius statistics", use_textfsm=True)
        sanitized_output = sanitize(sanitized_output, output)
        output = net_connect.send_command("show ap summary", use_textfsm=True)
        sanitized_output = sanitize(sanitized_output, output)
        output = net_connect.send_command("show fabric ap summary", use_textfsm=True)
        sanitized_output = sanitize(sanitized_output, output)
        output = net_connect.send_command("show wireless fabric client summary | i Number of", use_textfsm=True)
        sanitized_output = sanitize(sanitized_output, output)
        output = net_connect.send_command("show wireless client summary | i Number of", use_textfsm=True)
        sanitized_output = sanitize(sanitized_output, output)

        output = net_connect.send_command("show ap dot11 5 summary", use_textfsm=True)
        sanitized_output = count_radios_up(output, sanitized_output)
        output = net_connect.send_command("show wireless fabric summary", use_textfsm=True)
        sanitized_output = compute_fabric_state(output, sanitized_output)
        output = net_connect.send_command("show proc cpu platform", use_textfsm=True)
        logging.warning(f'proc cpu is {output}')
        sanitized_output = sanitize(sanitized_output, output)


        net_connect.disconnect()
        logging.warning(f"Out 2 is {sanitized_output}")
        return sanitized_output

cisco = {
'device_type': 'cisco_ios',
'host': '10.16.1.5',
'username': 'noc',
'password': ''
}

api.add_resource(GetCLI, '/metrics')
app.run(debug=True, port=4000, host="0.0.0.0", threaded=True, use_reloader=True)

