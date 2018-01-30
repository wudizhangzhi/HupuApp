#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 上午11:40
# @Author  : wudizhangzhi

import sys

sys.path.append('..')
from hupu.messages.messages import ScoreBoard, SocketMessage, BaseMessage
from hupu.messages.entries import StringEntry, IntEntry, ListEntry


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


class Test(BaseMessage):
    test = ListEntry('result.casino_update.desc')


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


def test_listentry():
    data = {"name": "wall", "args": [{"room": "NBA_PLAYBYPLAY_CASINO", "gid": "153733", "status": "2", "result": {
        "casino_update": [
            {"casino_id": 412051, "desc": "470人参与，394人选「能」", "max_bet": 0, "status": {"id": 1, "desc": "竞猜中"},
             "user_count": 470, "user_win_coins": 0},
            {"casino_id": 412050, "desc": "571人参与，469人选「能」", "max_bet": 0, "status": {"id": 1, "desc": "竞猜中"},
             "user_count": 571, "user_win_coins": 0},
            {"casino_id": 412046, "desc": "574人参与，303人选「能」", "max_bet": 0, "status": {"id": 1, "desc": "竞猜中"},
             "user_count": 574, "user_win_coins": 0},
            {"casino_id": 412041, "desc": "1014人参与，678人选「骑士」", "max_bet": 0, "status": {"id": 1, "desc": "竞猜中"},
             "user_count": 1014, "user_win_coins": 0},
            {"casino_id": 412037, "desc": "950人参与，604人选「不会」", "max_bet": 0, "status": {"id": 1, "desc": "竞猜中"},
             "user_count": 950, "user_win_coins": 0}]}, "online": "43万"}]}
    test = Test(data['args'][0])
    print(test.__dict__)
    print(test.test)


if __name__ == '__main__':
    # test_entry()
    test_listentry()