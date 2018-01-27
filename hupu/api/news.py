#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi


import time

from api import logger
from api.base import Base
from messages.messages import News, NewsDetail

log = logger.getLogger(__name__)


class NewsMixin(Base):
    def _getNews(self):
        """
        获取新闻
        """
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/nba/getNews'.format(self.api_version),
            params={
                'crt': int(time.time() * 1000),
                'night': 0,
                'channel': 'myapp',
                'client': self.client,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
                'entrance': '-1',
            }
        )

    def getNews(self):
        r = self._getNews().json()
        log.debug('获取新闻: {}'.format(r))
        return [News(news) for news in r['result']['data']]

    def _getNewsDetailSchema(self, nid):
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/nba/getNewsDetailSchema'.format(self.api_version),
            params={
                'crt': int(time.time() * 1000),
                'night': 0,
                'channel': 'myapp',
                'nid': nid,
                'nopic': 0,
                'catetype': 'news',
                'token': '',
                'top_ncid': -1,
                'replies': '',
                'client': self.client,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
                'entrance': '-1',
            }
        )

    def getNewsDetailSchema(self, nid):
        return NewsDetail(self._getNewsDetailSchema(nid).json()['result'])


if __name__ == '__main__':
    n = NewsMixin()
    print(n.getNewsDetailSchema(2255215))
