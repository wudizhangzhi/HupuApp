#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
from __future__ import absolute_import
from __future__ import print_function

import time

import colored
import requests

from hupu.api import logger
from hupu.utils import colored_text, parse_message
from hupu.messages.entries import to_text

log = logger.getLogger(__name__)
log = logger.getLogger('websocket')
import websocket


class HupuLiveWebSocket(object):
    def __init__(self, client, game, host=None, port=None, livetype=None):
        """
        虎扑文字直播websocket的基类
        :param client: 
        :param host: 
        :param port: 
        :param livetype: 文字直播的类型: [NBA, CBA]
        """
        self.client = client
        self.game = game
        self.pid = None
        self.HOST = host or '61.174.11.224'
        self.PORT = port or 3081

        self._connected = False  # 是否连接过

        if not livetype:
            self.livetype = 'NBA'
        else:
            self.livetype = livetype

        self.heart_beat_count = 0
        # self.ws = None

    def get_token(self):
        """
        post /socket.io/1/?client=008796750504411&t=1516672901&type=1&background=false
        response 包含 012tC3NMo1FlJsPhaSdh:50:60:websocket,htmlfile,xhr-polling,jsonp-polling
        :return: str example 012tC3NMo1FlJsPhaSdh
        """
        t = int(time.time())
        url = 'http://{host}:{port}/socket.io/1/?client={client}&t={t}&type=1&background=false'.format(host=self.HOST,
                                                                                                       port=self.PORT,
                                                                                                       client=self.client,
                                                                                                       t=t)
        r = requests.post(url, timeout=10)
        return r.text.split(':50:60')[0]

    @property
    def msg_send_to_get_match(self):
        return '5::/nba_v1:{"args":[{"room":"%s_HOME"}],"name":"join"}' % self.livetype

    def run(self):
        token = self.get_token()
        t = int(time.time())
        ws = websocket.WebSocketApp(
            "ws://{host}:{port}/socket.io/1/websocket/{token}/?client={client}&t={t}&type=1&background=false".format(
                host=self.HOST, port=self.PORT, token=token, client=self.client, t=t
            ),
            on_message=self.on_message,
            on_error=self.on_error,
            on_close=self.on_close,
            on_open=self.on_open,
        )
        log.debug('=== start websocket ===')
        ws.run_forever()

    def send(self, ws, message):
        ws.send(message)
        log.debug('send: {}'.format(message))

    def on_message(self, ws, message):
        """
        收到消息时候的回调
        """
        message = to_text(message)
        log.debug('receive: {}'.format(to_text(message)))
        msg = ''  # 返回的消息
        if message == '1::':
            msg = '2:::'
        elif message == '2::':
            if self._connected:
                msg = '2::'
            else:
                self._connected = True
                msg = '1::/nba_v1'
        elif message in ['1::/nba_v1']:  # 开始获取数据
            msg = self.on_match_message(ws, message)
        else:  # 数据部分
            socket_message = parse_message(message)
            if not socket_message:
                return
            if socket_message.room == 'NBA_HOME':
                pass
            else:
                self.on_live_message(ws, socket_message)
                self.heart_beat(ws)

            if socket_message.room_live_type == -1:  # 比赛结束
                print('----- 直播结束了, 即将退回菜单 -----', end='\n\r')
                time.sleep(3)
                ws.close()

        if msg:
            self.send(ws, msg)

    def on_match_message(self, ws, message):
        """
        返回比赛信息(比赛的gid, 比分, 情况)之后的回调
        例如 发送 '5::/nba_v1:{"args":[{"roomid":-1,"gid":100721,"pid":429,"room":"NBA_PLAYBYPLAY_CASINO"}],"name":"join"}'
        """
        # TODO 不太明白
        return '5::/nba_v1:{"args":[{"roomid":-1,"gid":%s,"pid":%s,"room":"NBA_PLAYBYPLAY_CASINO"}],"name":"join"}' % (
            self.game.gid, self.pid or 617)

    def on_live_message(self, ws, socket_message):
        """
        比赛直播信息
        :param ws: 
        :param message: 
        :return: 
        """
        pass

    def on_error(self, ws, error):
        # print('error: {}'.format(error))
        log.error('=== onerror: {} ==='.format(error))

    def on_close(self, ws):
        print('\n\r')
        print('|文字直播关闭|\n\r')
        log.debug('=== on close ===')

    def on_open(self, ws, *args, **kwargs):
        print('\n\r')
        print('|直播室连接中...|\n\r')
        log.debug('=== on open ===')

    def heart_beat(self, ws):
        """
        每隔5次心跳返回
        """
        self.heart_beat_count += 1
        if self.heart_beat_count > 5:
            self.heart_beat_count = 0
            heart_beat_msg = '2:::'
            self.send(ws, heart_beat_msg)
            log.debug('--- heart beat ---')


class HupuSocket(HupuLiveWebSocket):
    last_time = None  # 上次消息的时间
    last_pid = None  # 上次的pid

    def on_live_message(self, ws, socket_message):
        scoreboard, msgs = socket_message.scoreboard, socket_message.livemessges
        for msg in msgs:
            if not self.last_time or msg.t > self.last_time:
                self.print_live(scoreboard, msg)
                self.last_time = msg.t
        if hasattr(scoreboard, 'pid'):
            self.pid = scoreboard.pid
        return scoreboard, msgs

    def print_live(self, scoreboard, msg):
        """
        打印直播信息
        """
        print("{} \n\r".format(' | '.join((self.colored_scoreboard(scoreboard), str(msg)))))

    def colored_scoreboard(self, scoreboard):
        home_score = scoreboard.home_score
        away_score = scoreboard.away_score
        try:
            if int(home_score) > int(away_score):
                home_score = colored_text(home_score, colored.fg("red") + colored.attr("bold"))
            else:
                away_score = colored_text(away_score, colored.fg("red") + colored.attr("bold"))
        except Exception as e:
            pass
        text = '{home} {home_score}:{away_score} {away}  {process}'.format(
            home_score=home_score,
            away_score=away_score,
            process=scoreboard.process,
            home=self.game.home_name,
            away=self.game.away_name,
        )
        return text


def test_color():
    from messages.messages import ScoreBoard, Game
    game = Game({'gid': 153735, 'home_name': '勇士', 'away_name': '骑士'})

    hlws = HupuSocket(game=game, client='008796750504411', host='61.174.11.224', port=3081)
    sb = ScoreBoard({'home_score': 120, 'away_score': 89})
    # print(hlws.colored_scoreboard(sb))
    hlws.print_live(sb, '主持人: 雷阿伦左侧开球！this is a test!!!')


def test():
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
        i += 1
        time.sleep(1)
    #


if __name__ == '__main__':
    # hlws = HupuLiveWebSocket(client='008796750504411', host='61.174.11.224', port=3081)
    # hlws = HupuSocket(gid=153735, client='008796750504411', host='61.174.11.224', port=3081)
    # hlws.run()

    # test_color()
    test()
