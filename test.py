# -*- coding: utf-8 -*-

import md5
import time
import requests


class Tester(object):
    host = "127.0.0.1"
    port = "8908"
    url = "/service/search"

    def test(self, input_url=None):
        url = input_url or self.url
        request_url = "http://%s:%s%s" % (self.host, self.port, url)
        param_dict = {
            "level": "3",
            "signTime": str(int(time.time())),
            "deviceId": "A3645E9D8",
            "ptf": "M207",
        }
        token_input = param_dict["deviceId"] + param_dict["signTime"] + param_dict["level"] + param_dict["ptf"]
        m1 = md5.new()
        m1.update(token_input.encode(encoding='utf-8'))
        token = m1.hexdigest()
        param_dict["token"] = token
        param_dict["searchWord"] = "Timber"
        param = "token=%(token)s&level=%(level)s&signTime=%(signTime)s&deviceId=%(deviceId)s&ptf=%(ptf)s&searchWord=%(searchWord)s" % param_dict
        req_url = "%s?%s" % (request_url, param)
        print(req_url)
        response = requests.get(req_url)
        print(response.content)


class LocalTester(Tester):
    def __init__(self):
        self.host = "127.0.0.1"
        self.port = "8908"
        self.url = "/service/search"


class PreviewTester(Tester):
    def __init__(self):
        self.host = "10.98.16.215"
        self.port = "8908"
        self.url = "/service/search"


class ProductionTester(Tester):
    def __init__(self):
        self.host = "media.ptmi.gitv.ai"
        self.port = "8908"
        self.url = "/service/search"


if __name__ == '__main__':
    tester = LocalTester()
    tester.test()
