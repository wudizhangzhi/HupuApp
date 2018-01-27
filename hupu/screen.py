#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
"""
虎扑文字直播显示

---------------------------------
虎扑 Proudly presented by JRs.

帮助信息:
     j         Down           下移
     k         Up             上移
     space     Enter          进入
     q         Quit           退出


今日比赛:

     国王 105 vs 99 魔术 已结束
 ->  篮网 108 vs 109 雷霆 已结束
     骑士 102 vs 114 马刺 已结束
     尼克斯 112 vs 123 勇士 已结束
     凯尔特人 107 vs 108 湖人 已结束

----------------------------------
"""
import curses
import os
import time

from six import text_type

from api import logger
from utils import purge_text, text_to_list
from terminalsize import get_terminal_size

log = logger.getLogger(__name__)

shortcut = [
    ['j    ', 'Down      ', '下移'],
    ['k    ', 'Up        ', '上移'],
    ['space', 'Enter     ', '进入'],
    ['q    ', 'Quit      ', '退出'],
]


class Screen(object):
    def __init__(self, hupuapp, seq=(), **kwargs):
        self.hupuapp = hupuapp
        # self._t_columns, self._t_lines = os.get_terminal_size()
        self._t_columns, self._t_lines = get_terminal_size()

        self.title = '虎扑 Proudly presented by JRs.'
        self.mode_title = '今日比赛:'  # 模式
        self.mode = 'live'

        self._arrow_index = 0  # 箭头位置
        self._screen_lines = seq  # -屏幕显示的内容-, 数据部分
        self.screen = curses.initscr()
        self.screen.keypad(True)
        curses.noecho()
        curses.start_color()
        self._init_color()
        if kwargs:
            self.__dict__.update(kwargs)

    def __getitem__(self, item):
        return self._screen_lines[item]

    def __repr__(self):
        return '{!r}'.format(self._screen_lines)

    def _init_color(self):
        curses.use_default_colors()
        curses.init_pair(1, curses.COLOR_GREEN, -1)
        curses.init_pair(2, curses.COLOR_CYAN, -1)
        curses.init_pair(3, curses.COLOR_RED, -1)
        curses.init_pair(4, curses.COLOR_YELLOW, -1)

    @property
    def helper(self):
        title = '帮助信息:'
        title_length = len(title)
        return ['', title] + list(map(lambda x: ' ' * title_length + '     '.join(x), shortcut)) + ['']

    def _add_arrow(self, i, line):
        '''
        添加箭头
        :param i: 索引
        :param line: 行内容
        :return:
        '''
        arrow = ' ->  '
        if not isinstance(line, text_type):
            line = str(line)
        if i == self._arrow_index:
            line = arrow + line
        else:
            line = ' ' * len(arrow) + line
        return line

    def move_down(self):
        self._arrow_index += 1
        self._arrow_index = min(len(self._screen_lines) - 1, self._arrow_index)
        self.display()

    def move_up(self):
        self._arrow_index -= 1
        self._arrow_index = max(0, self._arrow_index)
        self.display()

    def empty(self):
        self._screen_lines = []

    def display(self, lines=None):
        """
        根据不同的模式显示内容
        :param lines: 打印的内容
        :return:
        """
        if lines:
            for line in lines:
                print('{}\n\r'.format(str(line)))
        else:
            if self.mode == 'live':
                self.build_games_menu()
            elif self.mode == 'news':
                self.build_news_menu()
            elif self.mode == 'newsdetail':
                self.build_news_detail()

    def set_screen(self, lines):
        self._screen_lines = lines

    def append(self, p_object):
        self._screen_lines.append(p_object)

    def set_mode(self, mode):
        self._arrow_index = 0  # 箭头还原
        if mode == 'live':
            self.mode_title = '今日比赛:'
            # TODO 或许还有其他操作
        elif mode == 'news':
            self.mode_title = '今日新闻:'
        elif mode == 'newsdetail':
            pass

        self.mode = mode
        self.display()

    def _build_static_part(self):
        # 标题区域
        self.screen.addstr(0, 0, self.title, curses.color_pair(3))
        # 帮助信息区域
        for i, helper in enumerate(self.helper):
            self.screen.addstr(1 + i, 0, helper, curses.color_pair(1))

        self.screen.addstr(len(self.helper) + 1, 0, self.mode_title)

    def build_games_menu(self):
        """
        构建比赛列表菜单
        """
        self.screen.clear()
        # 标题区域
        self.screen.addstr(0, 0, self.title, curses.color_pair(3))
        # 帮助信息区域
        for i, helper in enumerate(self.helper):
            self.screen.addstr(1 + i, 0, helper, curses.color_pair(1))

        self.screen.addstr(len(self.helper) + 1, 0, self.mode_title)

        start_index = len(self.helper) + 2

        # 比赛列表
        for i, line in enumerate(self._screen_lines):
            if i == self._arrow_index:
                self.screen.addstr(i + start_index, 0, self._add_arrow(i, line), curses.A_REVERSE)
            else:
                self.screen.addstr(i + start_index, 0, self._add_arrow(i, line))
        self.screen.refresh()

    def build_news_menu(self):
        """
        构建新闻菜单
        """
        self.screen.clear()
        self._build_static_part()

        start_index = len(self.helper) + 2
        # 新闻列表
        # 获取可视范围
        line_max = self._t_lines - start_index - 4
        show_end = self._arrow_index + 1 if self._arrow_index >= line_max else line_max
        show_start = show_end - line_max if show_end > line_max else 0

        for i, line in enumerate(self._screen_lines[show_start:show_end]):
            if i + show_start == self._arrow_index:
                self.screen.addstr(i + start_index, 0, self._add_arrow(i + show_start, line), curses.A_REVERSE)
                self.screen.addstr(i + start_index + 1, 0,
                                   self._add_arrow(i + show_start + 1,
                                                   '  ' + line.statistics_text),
                                   curses.color_pair(2))
            elif i > self._arrow_index:
                self.screen.addstr(i + start_index + 3, 0, self._add_arrow(i, line))
            else:
                self.screen.addstr(i + start_index, 0, self._add_arrow(i, line))
        self.screen.refresh()

    def build_news_detail(self, newsdetail=None):
        # 第一行中间显示标题
        newsdetail = newsdetail or self.newsdetail
        title = newsdetail.title
        x = abs((self._t_columns - len(title)) // 2)
        self.screen.addstr(0, x, title[:self._t_columns], curses.color_pair(1))

        # 空一行
        start_index = 2
        # 获取可视范围
        line_max = self._t_lines - start_index - 4
        show_end = self._arrow_index + 1 if self._arrow_index >= line_max else line_max
        show_start = show_end - line_max if show_end > line_max else 0

        for i, line in enumerate(self._screen_lines[show_start:show_end]):
            if i + show_start == self._arrow_index:
                self.screen.addstr(i + start_index, 0, line, curses.A_REVERSE)
            else:
                self.screen.addstr(i + start_index, 0, line)
        self.screen.refresh()

    def choose_game(self, index):
        from hupulivewebsocket import HupuSocket
        game_selected = self._screen_lines[index]
        hs = HupuSocket(game=game_selected, client=self.client_id)
        # try:
        hs.run()
        # except KeyboardInterrupt:
        log.debug('文字直播停止')
        print('文字直播停止\n\r')
        time.sleep(1)
        self.display()

    def choose_news(self, index):
        news = self._screen_lines[index]
        newsdetail = self.hupuapp.getNewsDetailSchema(news.nid)
        self.newsdetail = newsdetail  # 设置具体新闻
        # 正文
        content = purge_text(newsdetail.content)
        self.set_screen(text_to_list(content, self._t_columns - 5))
        self.set_mode('newsdetail')

    @property
    def endmsg(self):
        return '（づ￣3￣）づ╭❤～ bye~bye~'

    def quit(self):
        y = (self._t_lines - 1) // 2
        x = (self._t_columns - len(self.endmsg)) // 2
        self.screen.clear()
        self.screen.addstr(y, x, self.endmsg, curses.color_pair(3))
        self.screen.refresh()
        time.sleep(2)
        curses.endwin()

    def listen(self):
        while True:
            x = self.screen.getch()
            if x in [ord('j'), curses.KEY_DOWN]:
                self.move_down()
            elif x in [ord('k'), curses.KEY_UP]:
                self.move_up()
            elif x == ord('q'):
                if self.mode == 'newsdetail':
                    news = self.hupuapp.getNews()
                    self.set_screen(news)
                    self.set_mode('news')
                else:
                    self.quit()
                    break
            elif x == ord(' '):
                self.screen.clear()
                self.screen.refresh()
                if self.mode == 'live':
                    self.choose_game(self._arrow_index)

                elif self.mode == 'news':
                    self.choose_news(self._arrow_index)
