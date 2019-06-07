from flask import Flask
from flask_restful import Resource, Api, reqparse
from flask_cors import CORS
from netmiko import ConnectHandler
import argparse
import requests, json
import sys

app = Flask(__name__)
CORS(app)
api = Api(app)


class GetCLI(Resource):
    def get(self):
        net_connect = ConnectHandler(**cisco)
        output = net_connect.send_command("show ap summary", use_textfsm=True)
        net_connect.disconnect()
        print(f"Output is {output}")        
        return output


cisco = {
'device_type': 'cisco_ios',
'host': '192.168.30.205',
'username': 'lab',
'password': 'lab'
}

api.add_resource(GetCLI, '/metrics')
app.run(host='0.0.0.0', port='4000', debug=True)


