#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi
"""Hupu.
    Proudly presented by Hupu JRs.

Usage:
    hupu [-m MODE] [-u USERNAME] [-p PASSWORD] [-a APIVERSION]
    hupu -h | --help
    hupu -v | --version

Tips:
    Please hit Ctrl-C on the keyborad when you want to interrupt the game live.

Options:
    -u --username    input username.
    -p --password    input password.
    -a --apiversion  api version.[default: 7.1.15]
    -m --mode        run mode.[default: live, available: live news teamranks...]
    -h --help        Show this help message and exit.
    -v --version     Show version.
"""
from __future__ import print_function
import docopt
# python2 curses addstr乱码问题
import locale

locale.setlocale(locale.LC_ALL, '')

__version__ = '1.0.5'
__author__ = 'wudizhangzhi'
__all__ = ['HupuApp', 'start']

from hupu.hupuapp import HupuApp


def start():
    arguments = docopt.docopt(__doc__, version='Hupu {}'.format(__version__))
    hupulive = HupuApp(**arguments)
    hupulive.run()


if __name__ == '__main__':
    start()
