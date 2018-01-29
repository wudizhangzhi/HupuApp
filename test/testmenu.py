#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/29 下午8:28
# @Author  : wudizhangzhi
import sys

sys.path.append('..')
# from hupu.hupuapp import HupuApp
from hupu.menus.HupuMenu import HupuMenu


def test():
    # hupuapp = HupuApp()
    menu = HupuMenu(None)

    def choose_game(menu):
        teamranks = menu.hupuapp.getDatas()
        menu.items = teamranks
        menu.subtitle = '球队数据排行:'
        menu.draw()

    menu.register(ord(' '), 'live', choose_game)

    # menu.draw()
    # menu.listen()
    menu.quit()
    print(menu.key_bind_events)


if __name__ == '__main__':
    test()
