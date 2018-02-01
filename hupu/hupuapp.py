#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
"""Hupu.
    Proudly presented by Hupu JRs.

Usage:
    hupu [-m MODE] [-a APIVERSION] [-d DATATYPE] [-u USERNAME] [-p PASSWORD]
    hupu -h | --help
    hupu -v | --version

Tips:
    Please hit Ctrl-C on the keyborad when you want to interrupt the game live.

Options:
    -u USERNAME --username=USERNAME         Input username.
    -p PASSWORD --password=PASSWORD         Input password.
    -a APIVERSION --apiversion=APIVERSION   Api version.[default: 7.1.15]
    -m MODE --mode=MODE                     Run mode.Available: live news teamranks...[default: live]
    -d DATATYPE --datatype=DATATYPE         Player data type.Available: regular, injury, daily[default:regular]
    -h --help                               Show this help message and exit.
    -v --version                            Show version.
"""
from __future__ import print_function
# python2 curses addstr乱码问题
import locale

locale.setlocale(locale.LC_ALL, '')
import sys
import curses
import colored
import docopt
import traceback

from hupu.api.live import LiveMinxin
from hupu.api.login import LoginMixin
from hupu.api.news import NewsMixin
from hupu.api.datas import DatasMixin
from hupu.utils import colored_text, SYSTEM
from hupu.api import logger
from hupu.menus.HupuMenu import HupuMenu
from hupu.version import version

# if SYSTEM.lower() == 'windows':
#     # TODO debug window
#     reload(sys)
#     sys.setdefaultencoding('utf-8')
#     sys.stdout.encoding = 'cp65001'

log = logger.getLogger(__name__)

MODE_LIST = ['live', 'news', 'teamranks', 'playerdata']


class HupuApp(LiveMinxin, NewsMixin, LoginMixin, DatasMixin):
    def run(self):
        # 判断参数, 执行哪一种场景
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


def start():
    arguments = docopt.docopt(__doc__, version='Hupu {}'.format(version))
    # 处理参数
    arguments = {k.replace('--', ''): v for k, v in arguments.items()}
    hupulive = HupuApp(**arguments)
    hupulive.run()
