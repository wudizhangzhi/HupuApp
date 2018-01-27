#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Time    : 2018/1/27 下午2:22
# @Author  : wudizhangzhi

from __future__ import print_function
# python2 curses addstr乱码问题
import locale
locale.setlocale(locale.LC_ALL, '')

__version__ = '1.0.4'
__author__ = 'wudizhangzhi'
__all__ = ['HupuApp', 'main']

from hupu.hupuapp import HupuApp, main
