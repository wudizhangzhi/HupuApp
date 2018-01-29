#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
"""Hupu.
    Proudly presented by Hupu JRs.

Usage:
    hupu [-m MODE] [-u USERNAME] [-p PASSWORD] [-a APIVERSION]
    hupu -h | --help
    hupu -v | --version

Tips:
    Please hit Ctrl-C on the keyborad when you want to interrupt the game live.

Options:
    -u --username    input username.
    -p --password    input password.
    -a --apiversion  api version.[default: 7.1.15]
    -m --mode        run mode.[default: live, available: live news...]
    -h --help        Show this help message and exit.
    -v --version     Show version.
"""
import docopt
import requests
import curses
import colored
import traceback

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
from hupu.screen import Screen
from hupu.api import logger

log = logger.getLogger(__name__)

MODE_LIST = ['live', 'news', 'teamranks']


class HupuApp(LiveMinxin, NewsMixin, LoginMixin, DatasMixin):
    def run(self):
        # 判断参数, 执行哪一种场景
        # 1.没参数
        # 2.有用户名密码 -- 登录

        # 默认进入比赛文字直播模式
        mode = self._kwargs.get('MODE', '') or 'live'
        mode = mode.lower()
        assert mode in MODE_LIST, AttributeError('Expect mode is {}, got {}.'.format(', '.join(MODE_LIST), mode))
        try:
            screen = Screen(self, client_id=self.client)  # 显示的屏幕

            if mode == 'live':  # 文字直播模式
                games = self.getGames()
                screen.set_screen(games)

            elif mode == 'news':  # 新闻模式
                news = self.getNews()
                screen.set_screen(news)

            elif mode == 'teamranks':  # 球队数据模式
                teamranks = self.getDatas()
                screen.set_screen(teamranks)

            # 设置模式， 开始监听
            screen.set_mode(mode)
            screen.listen()

        except curses.error as e:
            curses.endwin()
            log.error(e)
            print(colored_text('窗口太小, 请调整窗口大小!', colored.fg("red") + colored.attr("bold")))
        except Exception as e:
            log.error(traceback.format_exc())
            if not curses.isendwin():
                curses.endwin()



def main():
    arguments = docopt.docopt(__doc__, version='Hupu 1.0')
    hupulive = HupuApp(**arguments)
    hupulive.run()


if __name__ == '__main__':
    main()
