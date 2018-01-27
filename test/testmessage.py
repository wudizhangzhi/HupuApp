#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午11:40
# @Author  : wudizhangzhi

import sys

sys.path.append('..')
from messages.messages import ScoreBoard, SocketMessage, BaseMessage
from messages.entries import StringEntry, IntEntry


class Integer:
    def __init__(self, name):
        self.name = name

    def __get__(self, instance, cls):
        if instance is None:
            return self
        else:
            return instance.__dict__[self.name]

    def __set__(self, instance, value):
        if not isinstance(value, int):
            raise TypeError('Expected an int')
        instance.__dict__[self.name] = value

    def __delete__(self, instance):
        del instance.__dict__[self.name]

def test_entry():
    # sb = ScoreBoard({'home_tid': 123})
    # print(sb.home_tid)
    # print(type(sb.home_tid))
    # print(type(sb))
    data = {
        "room": "NBA_PLAYBYPLAY_CASINO",
        "pid": 443,
        "result": {},
        "type": "nba",
        "gid": "153749",
        "room_live_type": "1"
    }

    class A(BaseMessage):
        __type__ = 'test'
        room_live_type = IntEntry('room_live_type')
        gid = ''

    a = A(data)
    print(a.__dict__)
    print(a.room_live_type, type(a.room_live_type))
    print(a.gid, type(a.gid))
    a.room_live_type = 12
    # s = SocketMessage(data)
    # print(s.__dict__)
    # print(s.room_live_type, type(s.room_live_type))


if __name__ == '__main__':
    test_entry()
