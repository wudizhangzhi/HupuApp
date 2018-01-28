#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi

import sys

sys.path.append('../..')
import time

from hupu.api import logger
from hupu.api.base import Base
from hupu.messages.messages import News, NewsDetail

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
        return [News(news) for news in r['result']['data'] if news.get('type') in News.news_type_list]

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
        j = self._getNewsDetailSchema(nid).json()
        log.debug(j)
        return NewsDetail(j['result'])

    def _getRecap7(self, gid):
        '''
        获取赛后
        :param gid: 
        :return: html
        '''
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/nba/getRecap7'.format(self.api_version),
            params={
                'client': self.client,
                'gid': gid,
                'nopic': 0,
                'night': 0,
                'entrance': '-1',
            }
        )


if __name__ == '__main__':
    n = NewsMixin()
    # print(n.getNewsDetailSchema(2255215))
    print(n._getRecap7(153764).text)
