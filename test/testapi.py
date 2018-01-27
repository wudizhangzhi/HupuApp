#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午4:16
# @Author  : wudizhangzhi

import sys
import locale
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



if __name__ == '__main__':
    # test_news_detail()
    test_unicode()
