#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:53
# @Author  : wudizhangzhi

import sys

sys.path.append('../..')
import time

from hupu.api import logger
from hupu.api.base import Base
from hupu.messages.messages import TeamRank

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
        )

    def getDatas(self):
        r = self._getData()
        return [TeamRank(i) for i in r.json()['data']['data']]

if __name__ == '__main__':
    d = DatasMixin()
    print(d.getDatas())
