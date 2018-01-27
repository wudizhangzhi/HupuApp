#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi
import json
import time

from hupu.api import logger
from hupu.api.base import Base
from hashlib import md5

log = logger.getLogger(__name__)


class LoginMixin(Base):
    def _login(self):
        username = self._kwargs.get('USERNAME', None)
        password = self._kwargs.get('PASSWORD', None)
        assert (username and password)
        url_login = 'https://games.mobileapi.hupu.com/1/{version}/user/loginUsernameEmail'.format(
            version=self.api_version)
        password_md5 = md5(password.encode('utf8')).hexdigest()
        # TODO 需要机械生成, 看上去调用两次可以登录
        # client = get_random_Imei()
        # client = '356359011002417'

        # client = '862561035110764'
        # client = '35672408061946'

        params = {
            'client': self.client
        }
        data = {
            'password': password_md5,
            'crt': str(int(time.time() * 1000)),
            'night': '0',
            'channel': 'oppo',
            'client': self.client,
            'time_zone': 'Asia/Shanghai',
            'android_id': self.android_id,
            'username': username,
        }
        log.debug('登录post信息: {}'.format(json.dumps(data, indent=4)))
        r = self.sess.post(url_login, params=params, data=data, timeout=10)
        r_json = r.json()
        log.debug('登录返回信息: {}'.format(r_json))
        login_success = self._is_login_success(r_json)
        if login_success:
            # 存储用户信息
            self._user_info = r_json['result']

    @staticmethod
    def _is_login_success(response_json):
        if 'error' in response_json:
            raise Exception('登录失败,返回: {}'.format(response_json['error']['text']))
        elif response_json['is_login'] == 1:
            print('登录成功')
            log.debug('登录成功')
            return True
        else:
            return False