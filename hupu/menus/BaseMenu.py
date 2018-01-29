#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/29 上午11:53
# @Author  : wudizhangzhi
# @File    : LiveMenu.py

# from cursesmenu import selection_menu
# from curtsies import FullscreenWindow, Input, FSArray
# from curtsies.fmtfuncs import red, bold, green, on_blue, yellow
# import curtsies.events
import curses

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


class BaseMenu(object):
    MAIN_PAGE = 0
    SUB_PAGE = 1

    def __init__(self, title=None, body_title=None, addition_title=None):
        self.title = title
        self.addition_title = addition_title
        self.body_title = body_title

        self.screen = None
        self.addition_items = list()
        self.items = list()
        self.sub_items = list()  # 下级item
        self.page_type = BaseMenu.MAIN_PAGE  # 页面类型

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
        # curses.init_pair(1, curses.COLOR_BLACK, curses.COLOR_WHITE)
        self.highlight = curses.A_REVERSE
        self.normal = curses.A_NORMAL

    def quit(self):
        curses.echo()
        curses.endwin()
        self.should_exit = True

    def draw(self):
        """
        Redraws the menu and refreshes the screen. Should be called whenever something changes that needs to be redrawn.
        """

        rows = 0  # 行数
        if self.title is not None:
            self.screen.addstr(rows, 2, self.title, curses.A_STANDOUT)
            rows += 2

        if self.addition_title and self.page_type == BaseMenu.MAIN_PAGE:
            self.screen.addstr(rows, 2, self.addition_title)
            rows += 1
            addition_start_row = rows
            for index, item in enumerate(self.addition_items):
                self.screen.addstr(addition_start_row + index, 4, item)
                rows += 1
            rows += 1

        if self.body_title is not None:
            self.screen.addstr(rows, 2, self.body_title, curses.A_BOLD)
            rows += 2

        screen_rows, screen_cols = self.screen.getmaxyx()
        row_max = screen_rows - rows - 1
        show_end = self.current_option + 1 if self.current_option >= row_max else row_max
        show_start = show_end - row_max if show_end > row_max else 0

        arrow = ' ->  '
        for index, item in enumerate(self.items[show_start:show_end]):
            if self.current_option == index + show_start:
                text_style = self.highlight
                self.screen.addstr(rows + index, 4, arrow + str(item), text_style)
            else:
                text_style = self.normal
                self.screen.addstr(rows + index, 4, ' ' * len(arrow) + str(item), text_style)
        self.screen.refresh()

    def move_up(self):
        if self.current_option > 0:
            self.current_option -= 1
        else:
            self.current_option = len(self.items) - 1
        self.draw()

    def move_down(self):
        if self.current_option < len(self.items) - 1:
            self.current_option += 1
        else:
            self.current_option = 0
        self.draw()

    def register(self, key, mode, func):
        self.key_bind_events[mode][key] = func

    def listen(self):
        while not self.should_exit:
            x = self.screen.getch()
            if x in [ord('j'), curses.KEY_DOWN]:
                self.move_down()
            elif x in [ord('k'), curses.KEY_UP]:
                self.move_up()
            elif x == ord('q'):
                if self.page_type == BaseMenu.MAIN_PAGE:
                    self.quit()
                else:
                    self.page_type = BaseMenu.MAIN_PAGE
                    self.current_option = 0
                    self.draw()
            elif x == ' ':
                pass
                # func = self.key_bind_events[self.mode][x]
                # func()


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
    base.listen()
