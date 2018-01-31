#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午4:16
# @Author  : wudizhangzhi

import sys
import locale
from logging.handlers import RotatingFileHandler

import os

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


def test_scoket_print():
    # from hupu.hupulivewebsocket import test
    # test()
    from hupu.messages.messages import Game
    # import sys
    import curses
    #
    def incurses(stdscr):
        stdscr.addstr(0, 0, "Exiting in ")
        stdscr.addstr(2, 0, "Hello World from Curses!")
        for i in range(5, -1, -1):
            stdscr.addstr(0, 11, str(i))
            stdscr.refresh()
            time.sleep(1)
        curses.endwin()

    curses.wrapper(incurses)
    # sys.stdout = os.fdopen(sys.stdout.fileno(), 'w', 0)
    sys.stdout = sys.__stdout__
    sys.stderr = sys.__stderr__
    print("After curses")
    # # curses.echo()
    # # curses.reset_shell_mode()
    # # print(sys.stdout)
    # game = Game({'gid': 153735, 'home_name': '勇士', 'away_name': '骑士'})
    # hlws = HupuSocket(game=game, client='008796750504411', host='127.0.0.1', port=5000)
    # hlws.run()
    # #
    # # def get_token():
    # #     return ''
    # #
    # # hlws.get_token = get_token
    # # hlws.run()
    i = 0
    while i < 100:
        print('this is a test. {}'.format(i))
        # sys.stdout.write('this is a test. {}'.format(i))
        i += 1
        time.sleep(1)


def test_init():
    hupuapp = HupuApp()
    print(hupuapp.getInit().json())


if __name__ == '__main__':
    # test_news_detail()
    # test_unicode()
    # test_logger()
    test_scoket_print()
    # test_init()
