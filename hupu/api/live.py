#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:58
# @Author  : wudizhangzhi
import datetime
import json
import time
from random import choice

from hupu.api import logger
from hupu.api.base import Base
from hupu.messages.messages import Game

log = logger.getLogger(__name__)


class LiveMinxin(Base):
    def _get_gamesobject_list(self, games_json):
        """
        :param games_json: 
        :return: 
        """
        today = datetime.datetime.today().strftime('%Y%m%d')

        games = games_json['result']['games']
        games_today = [i for i in games if i['day'] == today]
        games_today = games_today[0]
        # rank_url = games_today['rank_url']
        game_datas = games_today['data']

        game_object_list = []
        for _game_data in game_datas:
            gid = _game_data['gid']
            if len(gid) > 6:
                continue
            g = Game(_game_data)
            game_object_list.append(g)
        return game_object_list

    def _getGames(self, gametype=None):
        """
        获取比赛列表
        """
        if gametype:
            assert gametype.lower() in ['nba', 'cba']
        else:
            gametype = 'nba'
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/{}/getGames'.format(self.api_version, gametype),
            params={
                'crt': int(time.time() * 1000),
                'night': 0,
                'channel': 'myapp',
                'client': self.client,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
            }
        )

    def getGames(self, gametype=None):
        r = self._getGames(gametype).json()
        log.debug('获取比赛列表: {}'.format(r))
        return self._get_gamesobject_list(r)

    def getPlaybyplay(self, gid):
        """
        获取比赛直播信息
        :param gid: 
        :return: 
        """
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/room/getPlaybyplay'.format(self.api_version),
            params={
                'gid': gid,
                'crt': int(time.time() * 1000),
                'lid': '1',
                'roomid': '-1',
                'tag': 'nba',  # 默认nba
                'night': 0,
                'channel': 'myapp',
                'client': self.client,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
                'entrance': '-1',
            },
        )

    def getInit(self):
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/status/init'.format(self.api_version),
            params={
                'dv': '5.7.79',
                'crt': int(time.time() * 1000),
                'tag': 'nba',  # 默认nba
                'night': 0,
                'channel': 'myapp',
                'client': self.client,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
            },
        )

    def getIpAdress(self):
        host = port = None

        try:
            r_json = self.getInit().json()
            log.debug('ip地址获取: {}'.format(json.dumps(r_json, indent=2)))
            ip_adress_list = choice(r_json['result']['redirector'])
            log.debug('ip地址使用: {}'.format(ip_adress_list))
            tmp = ip_adress_list.split(':')
            host = tmp[0]
            port = int(tmp[1])
        except:
            pass
        return host, port
