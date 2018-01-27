#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午3:58
# @Author  : wudizhangzhi
import datetime
import time

from . import logger
from .base import Base

log = logger.getLogger(__name__)


class ForumMinxin(Base):
    def getThreadsList(self, gid):
        """
        TODO 
        """
        return self.sess.get(
            url='https://bbs.mobileapi.hupu.com/1/{}/recommend/getPlaybyplay'.format(self.api_version),
            params={
                'nav': 'nba,fifa,kog,follow,fitness,stylish,lol,pubg,cba,csl,lrw,hbl,epl,liga,seri,bund,fran,chlg,uefael,worldpre,national,afccl,zxb,confederationscup,clo,cl3,po1,nl1,tr1,sc1,ru1,be1,gr1,c1,kr1,ts1,no1,se1,ukr1,pl1,dk1,isr1,bu1,a1,ro1,gb2,gb3,gb4,es2,es3,l2,l3,rl,it2,fr2,po2,nl2,bra1,ar1n,jap1,jap2,jap3,rsk1,rsk2,aus1,tha1,hgkg,sa1,uae1,qsl,uz1,irn1,mys1,inil,sin1,mya1,in1l',
                'clientId': self.client,
                'client': '26439107',
                'crt': int(time.time() * 1000),
                'night': 0,
                'lastTid': 0,
                'stamp': int(time.time()),
                'isHome': 1,
                'time_zone': 'Asia/Shanghai',
                'android_id': self.android_id,
                'additionTid': '',
                'unfollowTid': '',
            },
        )
