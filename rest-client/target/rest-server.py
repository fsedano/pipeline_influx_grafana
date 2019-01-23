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

def post_util(ap, tag):
        wlc = WLC("10.51.65.154")
        try:
            ret = wlc.move_ap(ap, tag)
        except ConnectionError:
            ret = {'failed': True}
        return ret

class startAP1(Resource):
    def post(self):
        ap = "00:a6:ca:6c:d5:90"
        tag = "my_tag1"
        print("Start AP1: {}".format(ap))
        return post_util(ap, tag)

class stopAP1(Resource):
    def post(self):
        ap = "00:a6:ca:6c:d5:90"
        tag = "tag2"
        print("Stop AP1: {}".format(ap))
        return post_util(ap, tag):

class startAP2(Resource):
    def post(self):
        ap = "f4:4e:05:43:34:54"
        tag = "my_tag1"
        print("Start AP1: {}".format(ap))
        return post_util(ap, tag)

class stopAP2(Resource):
    def post(self):
        ap = "f4:4e:05:43:34:54"
        tag = "tag2"
        print("Stop AP1: {}".format(ap))
        return post_util(ap, tag)

api.add_resource(startAP1, '/start_ap1')
api.add_resource(stopAP1, '/stop_ap1')
api.add_resource(startAP2, '/start_ap2')
api.add_resource(stopAP2, '/stop_ap2')

app.run(host='0.0.0.0', port='4000', debug=True)


