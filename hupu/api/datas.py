#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi
import json
import sys

sys.path.append('../..')
import time

from hupu.api import logger
from hupu.api.base import Base
from hupu.messages.messages import TeamRank, PlayData

log = logger.getLogger(__name__)


class DatasMixin(Base):
    def _getData(self):
        """
        获取球队数据排行
        """
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/data/nba'.format(self.api_version),
            params={
                'client': self.client,
                'offline': 'json',
                'webp': 1,
            }
        ).json()

    def getDatas(self):
        r_json = self._getData()
        log.debug(json.dumps(r_json, indent=2))
        return [TeamRank(i) for i in r_json['data']['data']]

    def _getPlayerDataInGenernal(self, datatype='regular'):
        """
        球员数据接口
        :param datatype: regular, injury, daily
        """
        return self.sess.get(
            url='https://games.mobileapi.hupu.com/1/{}/nba/getPlayerDataInGeneral/'.format(self.api_version),
            params={
                'client': self.client,
                'offline': 'json',
                'type': datatype,
                'webp': 1,
            }
        ).json()

    def getPlayerDataInGenernal(self, datatype='regular'):
        r_json = self._getPlayerDataInGenernal(datatype)
        log.debug(log.debug(json.dumps(r_json, indent=2)))
        return [PlayData(i) for i in r_json['data']]


if __name__ == '__main__':
    d = DatasMixin()
    print(d.getDatas())
