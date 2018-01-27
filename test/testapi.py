#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午4:16
# @Author  : wudizhangzhi

import sys
import locale
from logging.handlers import RotatingFileHandler

locale.setlocale(locale.LC_ALL, '')
import time
sys.path.append('..')


def test_news_detail():
    from news import NewsMixin
    n = NewsMixin()
    print(n.getNewsDetailSchema(2255215))


from hupu.hupuapp import HupuApp


def test_unicode():

    hupu = HupuApp()
    games = hupu.getGames()
    print(games)

    import curses
    win = curses.initscr()
    win.clear()
    win.addstr(0, 0, str('   我们在一起'))
    win.refresh()
    time.sleep(1)
    curses.endwin()

def test_logger():
    import logging
    log = logging.getLogger('websocket')
    fh = RotatingFileHandler('log.log')
    fh.setLevel(logging.DEBUG)

    fh.setFormatter(logging.Formatter(
        '%(asctime)s - %(levelname)s - %(name)s:%(lineno)s: %(message)s'))
    log.addHandler(fh)
    # logger.setLevel(logging.CRITICAL)
    # logger.setLevel(logging.CRITICAL)
    # logger.setLevel(logging.CRITICAL)
    # logger.setLevel(logging.CRITICAL)
    import websocket
    from websocket import _logging

    _logging.error('hello')

if __name__ == '__main__':
    # test_news_detail()
    # test_unicode()
    test_logger()