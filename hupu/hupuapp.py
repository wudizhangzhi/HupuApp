#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi

import requests
import curses
import colored
import traceback

import sys

try:
    # 不打印ssl警告
    requests.packages.urllib3.disable_warnings()
except ImportError:
    pass

from hupu.api.live import LiveMinxin
from hupu.api.login import LoginMixin
from hupu.api.news import NewsMixin
from hupu.api.datas import DatasMixin
from hupu.utils import colored_text
from hupu.api import logger
from hupu.menus.HupuMenu import HupuMenu

log = logger.getLogger(__name__)

MODE_LIST = ['live', 'news', 'teamranks', 'playerdata']


class HupuApp(LiveMinxin, NewsMixin, LoginMixin, DatasMixin):
    def run(self):
        # 判断参数, 执行哪一种场景
        # 1.没参数
        # 2.有用户名密码 -- 登录

        # 默认进入比赛文字直播模式
        mode = self._kwargs.get('mode', '') or 'live'
        mode = mode.lower()
        assert mode in MODE_LIST, AttributeError('Expect mode is {}, got {}.'.format(', '.join(MODE_LIST), mode))
        try:
            hupumenu = HupuMenu(self)
            items = []
            if mode == 'live':  # 文字直播模式
                items = self.getGames()

            elif mode == 'news':  # 新闻模式
                items = self.getNews()
                hupumenu.body_title = '新闻:'

            elif mode == 'teamranks':  # 球队数据模式
                items = self.getDatas()
                hupumenu.body_title = '球队数据:'

            elif mode == 'playerdata':  # 球队数据模式
                datatype = self._kwargs.get('datatype')
                if not datatype or datatype not in ['regular', 'injury', 'daily']:
                    datatype = 'regular'
                datatype = datatype.lower()
                items = self.getPlayerDataInGenernal(datatype)
                hupumenu.body_title = '球员数据:'

            if not items:
                raise Exception('没有数据!')
            hupumenu.set_items(items)
            hupumenu.mode = mode

            hupumenu.draw()
            hupumenu.listen()

        except curses.error as e:
            curses.endwin()
            log.error(e)
            print(colored_text('窗口太小, 请调整窗口大小!', colored.fg("red") + colored.attr("bold")))
        except Exception as e:
            log.error(traceback.format_exc())
            if not curses.isendwin():
                curses.endwin()
            print(e)
