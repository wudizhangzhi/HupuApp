#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/29 上午11:53
# @Author  : wudizhangzhi
# @File    : LiveMenu.py

from __future__ import print_function

import traceback

from six import integer_types
import time
import curses
from collections import defaultdict

from hupu.api import logger
from hupu.utils import SYSTEM, PREFERREDENCODING

log = logger.getLogger(__name__)

"""
虎扑文字直播显示

---------------------------------
title

addition_title:
     addition_item
     addition_item
     addition_item
     addition_item


body_title:

     item
     item
     item
     item
     item
     item
----------------------------------
"""
key_bind_events = defaultdict(dict)

MAIN_PAGE = 0
SUB_PAGE = 1


def bind_event(key_list, mode_list=None):
    """
    绑定按键, 模块 和 回调
    :param key_list: 按键列表
    :param mode_list: 模式列表
    :return:
    """
    if not mode_list:
        mode_list = ['default']

    def decorator(func):
        for key in key_list:
            for mode in mode_list:
                if not isinstance(key, integer_types):
                    key = ord(key)
                key_bind_events[key][mode] = func
        return func

    return decorator


class BaseMenu(object):
    key_bind_events = defaultdict(dict)

    def __init__(self, title=None, body_title=None, addition_title=None):
        self.title = title
        self.addition_title = addition_title
        self.body_title = body_title
        self.body_title_buffer = body_title  # 缓存主体标题

        self.screen = None
        self.addition_items = list()
        self.items = list()
        self.items_buffer = list()  # 缓存item
        self.mode = 'live'
        self.page_type = MAIN_PAGE  # 页面类型

        self.should_exit = False
        self.highlight = None
        self.normal = None
        self.current_option = 0

        self._set_up_screen()
        self._set_up_colors()

    def set_items(self, items):
        self.items = items

    def _set_up_screen(self):
        self.screen = curses.initscr()
        self.screen.keypad(True)
        curses.noecho()
        curses.start_color()

    def _set_up_colors(self):
        curses.start_color()
        curses.use_default_colors()
        curses.init_pair(1, curses.COLOR_GREEN, -1)
        curses.init_pair(2, curses.COLOR_CYAN, -1)
        curses.init_pair(3, curses.COLOR_RED, -1)
        curses.init_pair(4, curses.COLOR_YELLOW, -1)
        curses.init_pair(5, curses.COLOR_WHITE, curses.COLOR_RED)
        curses.init_pair(6, curses.COLOR_WHITE, curses.COLOR_BLUE)
        self.highlight = curses.A_REVERSE
        self.normal = curses.A_NORMAL

    def reset_current_option(self):
        self.current_option = 0

    def quit(self):
        y, x = self.screen.getmaxyx()
        row = (y - 1) // 2
        col = (x - len(self.endmsg)) // 2
        self.screen.clear()
        self.addstr(row, col, self.endmsg, curses.color_pair(3))
        self.screen.refresh()
        time.sleep(2)

        curses.echo()
        curses.endwin()
        self.should_exit = True

    @property
    def endmsg(self):
        if SYSTEM == 'windows':
            return 'bye~bye~'
        return '（づ￣3￣）づ╭❤～ bye~bye~'

    def addstr(self, *args):
        if len(args) == 1:
            self.screen.addstr(args[0])
        else:
            content = args[2]
            if SYSTEM == 'windows':
                try:
                    content = args[2].decode('utf8').encode(PREFERREDENCODING)
                except Exception as e:
                    log.debug(traceback.format_exc())
            self.screen.addstr(args[0], args[1], content, *args[3:])

    def draw(self):
        """
        draw screen
        """
        self.clear_screen()
        rows = 0  # 行数
        if self.title is not None:
            self.addstr(rows, 2, self.title, curses.color_pair(5))
            rows += 2

        if self.addition_title and self.page_type == MAIN_PAGE:
            self.addstr(rows, 2, self.addition_title, curses.color_pair(6))
            rows += 1
            addition_start_row = rows
            for index, item in enumerate(self.addition_items):
                self.addstr(addition_start_row + index, 4, item)
                rows += 1
            rows += 1

        if self.body_title is not None:
            self.addstr(rows, 2, self.body_title, curses.A_BOLD)
            rows += 2

        screen_rows, screen_cols = self.screen.getmaxyx()
        row_max = screen_rows - rows - 1
        show_end = self.current_option + 1 if self.current_option >= row_max else row_max
        show_start = show_end - row_max if show_end > row_max else 0

        arrow = ' ->  '
        for index, item in enumerate(self.items[show_start:show_end]):
            if self.current_option == index + show_start:
                text_style = self.highlight
                self.addstr(rows + index, 4, arrow + str(item), text_style)
            else:
                text_style = self.normal
                self.addstr(rows + index, 4, ' ' * len(arrow) + str(item), text_style)
        self.screen.refresh()

    @bind_event(['k', curses.KEY_UP])
    def move_up(self):
        if self.current_option > 0:
            self.current_option -= 1
        else:
            self.current_option = len(self.items) - 1
        self.draw()

    @bind_event(['j', curses.KEY_DOWN])
    def move_down(self):
        if self.current_option < len(self.items) - 1:
            self.current_option += 1
        else:
            self.current_option = 0
        self.draw()

    @bind_event(['q', curses.KEY_EXIT])
    def back_or_quit(self):
        if self.page_type == MAIN_PAGE:
            self.quit()
        else:
            self.backto_mainpage()

    def clear_screen(self):
        """
        清除屏幕
        """
        self.screen.clear()

    def register(self, key, mode, func):
        key_bind_events[key][mode] = func

    def backto_mainpage(self):
        self.reset_current_option()
        self.page_type = MAIN_PAGE
        self.body_title, self.body_title_buffer = self.body_title_buffer, self.body_title
        self.items, self.items_buffer = self.items_buffer, self.items
        self.draw()

    def jumpto_subpage(self, title, subitems):
        self.reset_current_option()
        self.page_type = SUB_PAGE
        self.body_title, self.body_title_buffer = title, self.body_title
        self.items, self.items_buffer = subitems, self.items
        self.draw()

    def listen(self):
        while not self.should_exit:
            x = self.screen.getch()
            func = key_bind_events[x].get(self.mode)
            if not func:
                func = key_bind_events[x].get('default')
            if func:
                self.clear_screen()
                func(self)

if __name__ == '__main__':
    base = BaseMenu('虎扑 Proudly presented by JRs.', '今日比赛:')
    base.set_items([str(i) for i in range(10)])
    base.addition_title = '帮助信息:'
    base.addition_items = [
        'j         Down           下移',
        'k         Up             上移',
        'space     Enter          进入',
        'q         Quit           退出',
    ]
    base.draw()

    base.register(ord(' '), 'default', lambda x: print(x.title))
    base.listen()
