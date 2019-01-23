from flask import Flask
from flask_restful import Resource, Api, reqparse
from flask_cors import CORS

import argparse
import requests, json
import sys

app = Flask(__name__)
CORS(app)
api = Api(app)

class WLC:
    def __init__(self, controller_ip):
        self.controller_ip = controller_ip
        self.headers = {
            'Accept': "application/yang-data+json",
            'Content-Type': "application/yang-data+json",
            'Authorization': "Basic Y2lzY286Y2lzY28=",
            'cache-control': "no-cache"
        }
        self.url = f"https://{self.controller_ip}/restconf/data/Cisco-IOS-XE-wireless-ap-cfg:ap-cfg-data/ap-tags/ap-tag/"

    def _change_policy_tag_payload(self, payload, mac, new_policy_tag):
        _payload = json.loads(payload.text)
        print(_payload)
        for entry in _payload['Cisco-IOS-XE-wireless-ap-cfg:ap-tag']:
            if entry['ap-mac'] == mac:
                entry['policy-tag'] = new_policy_tag
        return json.dumps(_payload)

    def move_ap(self, mac, tag):
        payload = requests.request("GET", self.url, headers=self.headers, verify=False)   
        print("Got config from Controller:")
        print(payload.text)
        if True:
            print(f"Moving AP:{mac} to tag:{tag}")
            payload = self._change_policy_tag_payload(payload, mac, tag)
            response = requests.request("PATCH", self.url, data=payload, headers=self.headers, verify=False)
            print(response.text)
            print("Got config from Controller, after update:")
            response = requests.request("GET", self.url, headers=self.headers, verify=False)
            print(response.text)
        return {'ok':True}

class startAP1(Resource):
    def post(self):
        print("Start AP1")
        wlc = WLC("10.51.65.154")
        try:
            ret = wlc.move_ap("00:a6:ca:6c:d5:90", "tag2")
        except ConnectionError:
            ret = {'failed': True}
        return ret

api.add_resource(startAP1, '/start_ap1')

app.run(host='0.0.0.0', port='4000', debug=True)


