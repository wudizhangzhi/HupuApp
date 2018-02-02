#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/2/2 下午2:23
# @Author  : wudizhangzhi
# @File    : testsocketprint.py
import logging

import os

log = logging.basicConfig()
import websocket
import sys
import curses


class StdOutWrapper:
    text = ""

    def write(self, txt):
        self.text += txt
        self.text = '\n'.join(self.text.split('\n')[-30:])

    def get_text(self, beg, end):
        return '\n'.join(self.text.split('\n')[beg:end])


def test():
    """
    测试websocket app 建立连接后的print不正常
    :return:
    """
    global count
    count = 0

    def on_message(ws, message):
        global count
        print('receive: {}'.format(message))
        count += 1
        content = 'reply: {}'.format(count)
        ws.send(content)
        print('send: {}'.format(content))

    def on_error(ws, error):
        print('error: {}'.format(error))

    def on_close(ws):
        print('close')
        print('sys.stdout: {}'.format(sys.stdout))
        print('sys.__stdout__: {}'.format(sys.__stdout__))

    def on_open(ws, *args, **kwargs):
        print('open : {!r}  {!r}'.format(args, kwargs))
        ws.send('hello world')

    ########################
    stdscr = curses.initscr()
    curses.nocbreak()  # 关闭字符终端功能（只有回车时才发生终端）
    stdscr.keypad(0)
    curses.echo()
    curses.endwin()

    sys.stdout = os.fdopen(sys.stdout.fileno(), 'w', 0)

    print('sys.stdout: {}'.format(sys.stdout.name))
    print('sys.__stdout__: {}'.format(sys.__stdout__))

    ws = websocket.WebSocketApp(
        "ws://echo.websocket.org/",
        on_message=on_message,
        on_error=on_error,
        on_close=on_close,
        on_open=on_open,
    )
    ws.run_forever()

    print('-- closed --')

    # if not sys.stdout.closed:
        # sys.stdout = os.fdopen(sys.stdout.fileno(), 'w', 0)
        # pass
    print('reset sys.stdout')


if __name__ == '__main__':
    test()
