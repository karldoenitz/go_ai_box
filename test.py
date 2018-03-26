# -*- coding: utf-8 -*-

import md5
import time
import requests

url = "http://127.0.0.1:8908/service/search"
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
req_url = "%s?%s" % (url, param)
print(req_url)
response = requests.get(req_url)
print(response.content)
