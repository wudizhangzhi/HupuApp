#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
"""Hupu.
    Proudly presented by Hupu JRs.

Usage:
    hupu [-m MODE] [-a APIVERSION] [-d DATATYPE] [-u USERNAME] [-p PASSWORD]
    hupu -h | --help
    hupu -v | --version

Tips:
    Please hit Ctrl-C on the keyborad when you want to interrupt the game live.

Options:
    -u USERNAME --username=USERNAME         Input username.
    -p PASSWORD --password=PASSWORD         Input password.
    -a APIVERSION --apiversion=APIVERSION   Api version.[default: 7.1.15]
    -m MODE --mode=MODE                     Run mode.Available: live news teamranks...[default: live]
    -d DATATYPE --datatype=DATATYPE         Player data type.Available: regular, injury, daily[default:regular]
    -h --help                               Show this help message and exit.
    -v --version                            Show version.
"""
from __future__ import print_function
import docopt
# python2 curses addstr乱码问题
import locale

locale.setlocale(locale.LC_ALL, '')

__version__ = '1.0.7'
__author__ = 'wudizhangzhi'
__all__ = ['HupuApp', 'start']

from hupu.hupuapp import HupuApp


def start():
    arguments = docopt.docopt(__doc__, version='Hupu {}'.format(__version__))
    # 处理参数
    arguments = {k.replace('--', ''): v for k, v in arguments.items()}
    hupulive = HupuApp(**arguments)
    hupulive.run()


if __name__ == '__main__':
    start()
