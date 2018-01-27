#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi


import requests

try:
    requests.packages.urllib3.disable_warnings()
except ImportError:
    pass
from user_agent import generate_user_agent

from hupu.utils import getSortParam, get_android_id, get_random_Imei
from hupu.api import logger

log = logger.getLogger(__name__)

HUPU_API_VERSION = '7.1.15'

MODE_LIST = ['live', 'news']


class SignSession(requests.Session):  # 继承
    def request(self, method, url,
                params=None, data=None, headers=None, cookies=None, files=None,
                auth=None, timeout=None, allow_redirects=True, proxies=None,
                hooks=None, stream=None, verify=None, cert=None, json=None):
        if data:
            sign = getSortParam(**data)
            data.update({'sign': sign})
        if len(params) > 2:
            sign = getSortParam(**params)
            params.update({'sign': sign})
        return super(SignSession, self).request(method, url,
                                                params=params, data=data, headers=headers, cookies=cookies, files=files,
                                                auth=auth, timeout=timeout, allow_redirects=allow_redirects,
                                                proxies=proxies,
                                                hooks=hooks, stream=stream, verify=verify, cert=cert, json=json)


class Base(object):
    def __init__(self, **kwargs):
        self._kwargs = kwargs
        self.api_version = HUPU_API_VERSION
        # 初始化设备信息
        self.sess = self._init_session()

        self._user_info = {}  # 用于存储用户信息

    def _init_session(self):
        sess = SignSession()
        headers = {
            'User-Agent': generate_user_agent(os=('android',)) + ' kanqiu/{}.13305/7214 isp/-1 network/-1'.format(
                self.api_version),
            'Content-Type': 'application/x-www-form-urlencoded',
        }
        sess.headers = headers
        sess.verify = False  # 关闭ssl验证
        return sess

    @property
    def client(self):
        if not (hasattr(self, '_client') and self._client):
            setattr(self, '_client', get_random_Imei())
        return self._client

    @property
    def android_id(self):
        if not (hasattr(self, '_android_id') and self._android_id):
            setattr(self, '_android_id', get_android_id())
        return self._android_id