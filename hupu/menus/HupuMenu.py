#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/29 下午8:17
# @Author  : wudizhangzhi
import time

from hupu.menus.BaseMenu import BaseMenu, SUB_PAGE, bind_event
from hupu.api import logger
from hupu.messages.entries import to_text
from hupu.utils import purge_text, text_to_list
from hupu.hupulivewebsocket import HupuSocket

log = logger.getLogger(__name__)

HELPER_LINES = [
    'j         Down           下移',
    'k         Up             上移',
    'space     Enter          进入',
    'q         Quit           退出',
]


class HupuMenu(BaseMenu):
    def __init__(self, hupuapp, title=None, body_title=None, addition_title=None):
        super(HupuMenu, self).__init__(title, body_title, addition_title)
        self.hupuapp = hupuapp

        self.title = '虎扑 Proudly presented by JRs.'
        self.body_title = '今日比赛:'
        self.addition_title = '帮助信息:'
        self.addition_items = HELPER_LINES

    # def choose_game(self):
    #     self.items[self.current_option]

    @bind_event(' ', 'teamranks')
    def choose_teamranks(self):
        teamrank = self.items[self.current_option]
        self.jumpto_subpage(teamrank.title, teamrank.to_table)

    @bind_event(' ', 'news')
    def choose_news(self):
        news = self.items[self.current_option]
        newsdetail = self.hupuapp.getNewsDetailSchema(news.nid)
        # 正文
        content = purge_text(to_text(newsdetail.content))
        # 中文显示问题
        self.jumpto_subpage(newsdetail.title, text_to_list(content, self.screen.getmaxyx()[0]))

    @bind_event(' ', 'live')
    def choose_live(self):
        self.clear_screen()
        self.screen.refresh()
        game_selected = self.items[self.current_option]
        hs = HupuSocket(game=game_selected, client=self.hupuapp.client)
        # try:
        hs.run()
        # except KeyboardInterrupt:
        log.debug('文字直播停止')
        # print('文字直播停止\n\r')
        time.sleep(1)
        self.draw()
